# gspec

[![Go project version](https://badge.fury.io/go/github.com%2Fbroothie%2Fgspec.svg)](https://badge.fury.io/go/github.com%2Fbroothie%2Fgspec)
[![Go Report Card](https://goreportcard.com/badge/github.com/broothie/gspec)](https://goreportcard.com/report/github.com/broothie/gspec)
[![codecov](https://codecov.io/gh/broothie/gspec/branch/main/graph/badge.svg?token=6CLN4sDTk5)](https://codecov.io/gh/broothie/gspec)
[![gosec](https://github.com/broothie/gspec/actions/workflows/gosec.yml/badge.svg)](https://github.com/broothie/gspec/actions/workflows/gosec.yml)
[![GitHub](https://img.shields.io/github/license/broothie/gspec)](https://opensource.org/license/mit/)

`gspec` is a testing framework for Go inspired by Ruby's [RSpec](https://rspec.info/). It adds nested example groups, lazy per-case values, hooks, and fluent expectations while continuing to use Go's built-in `testing` package.

`gspec` requires Go 1.27 or later and has no third-party runtime dependencies.

## Installation

```shell
go get github.com/broothie/gspec
```

## Quick start

Matchers are designed to read naturally when the `match` package is dot-imported:

```go
package calculator_test

import (
	"testing"

	"github.com/broothie/gspec"
	. "github.com/broothie/gspec/match"
)

func TestAddition(t *testing.T) {
	gspec.Describe(t, "addition", func(c *gspec.Context) {
		c.It("returns the sum of its operands", func(c *gspec.Case) {
			c.Expect(1 + 2).To(Equal(3))
			c.Expect(1 + 2).NotTo(Equal(4))
		})
	})
}
```

`gspec.Describe` opens a labelled root group. Use `gspec.Run` when the root does not need a label. Each call to `It` becomes a normal Go subtest, so existing `go test` tooling continues to work.

## Groups

Use `Describe` to name a subject and `Context` to describe a condition. Groups can be nested arbitrarily and inherit lets and hooks from their parents.

```go
func TestGreeting(t *testing.T) {
	gspec.Run(t, func(c *gspec.Context) {
		c.Describe("Greeting", func(c *gspec.Context) {
			c.Context("when a name is present", func(c *gspec.Context) {
				c.It("includes the name", func(c *gspec.Case) {
					c.Expect("Hello, Gopher!").To(ContainSubstring("Gopher"))
				})
			})
		})
	})
}
```

## Lazy values

`Let` defines a type-safe value that is evaluated only when a case first calls `Get`. The result is cached for the rest of that case and evaluated independently for every other case.

Nested groups can override a let with `Set`. Lets may also depend on other lets:

```go
func TestGreeting(t *testing.T) {
	gspec.Describe(t, "Greeting", func(c *gspec.Context) {
		name := c.Let(func(c *gspec.Case) string { return "Gopher" })
		greeting := c.Let(func(c *gspec.Case) string {
			return "Hello, " + c.Get(name) + "!"
		})

		c.It("greets the default name", func(c *gspec.Case) {
			c.Expect(c.Get(greeting)).To(Equal("Hello, Gopher!"))
		})

		c.Context("with another name", func(c *gspec.Context) {
			c.Set(name, func(c *gspec.Case) string { return "Rubyist" })

			c.It("uses the overridden name", func(c *gspec.Case) {
				c.Expect(c.Get(greeting)).To(Equal("Hello, Rubyist!"))
			})
		})
	})
}
```

## Hooks

`BeforeEach` and `AfterEach` register setup and cleanup functions. Hooks are inherited by nested groups and receive the current `Case`, allowing them to access lets with `Get`.

```go
gspec.Run(t, func(c *gspec.Context) {
	values := c.Let(func(c *gspec.Case) *[]int { return &[]int{} })

	c.BeforeEach(func(c *gspec.Case) {
		caseValues := c.Get(values)
		*caseValues = append(*caseValues, 1)
	})

	c.AfterEach(func(c *gspec.Case) {
		*c.Get(values) = nil
	})

	c.It("runs between the hooks", func(c *gspec.Case) {
		c.Expect(*c.Get(values)).To(Contain(1))
	})
})
```

See the executable [hooks example](./examples/hooks_test.go) for a complete setup and cleanup flow.

## Expectations

Inside a case, `Expect` captures an actual value. `To` reports a failure when its matcher does not match, while `NotTo` reports a failure when it does.

```go
c.Expect(actual).To(Equal(expected))
c.Expect(actual).NotTo(Equal(unexpected))
```

Outside a `Case`, the package-level form reports to any compatible test value:

```go
gspec.Expect(t, actual).To(Equal(expected))
```

### Matchers

| Category | Matchers |
| --- | --- |
| Values | `Equal`, `BeNil`, `Satisfy` |
| Ordering | `BeLessThan`, `BeAtMost`, `BeGreaterThan`, `BeAtLeast`, `BeCloseTo` |
| Collections | `BeEmpty`, `HaveLength`, `Contain`, `ConsistOf` |
| Strings | `ContainSubstring`, `HavePrefix`, `HaveSuffix`, `MatchRegexp` |
| Errors | `HaveOccurred`, `BeError`, `BeErrorType` |
| Actions | `Change`, `Panic`, `PanicWith` |

Some matcher semantics are worth calling out:

- `Equal` performs deep equality, so it supports slices and maps.
- `Contain` requires every expected element to be present.
- `ConsistOf` ignores order, but duplicate elements remain significant.
- `BeError` uses `errors.Is`; `BeErrorType` uses `errors.As`.
- `Change` evaluates a value before and after running an action.
- `PanicWith` deeply compares the recovered panic value.

Most matcher types are inferred from their arguments. Matchers whose value type does not appear in an argument need an explicit type argument:

```go
c.Expect(pointer).To(BeNil[*Widget]())
c.Expect(values).To(BeEmpty[[]int]())
c.Expect(values).To(HaveLength[[]int](3))
c.Expect(err).To(BeErrorType[*ValidationError]())
```

For a one-off assertion, `Satisfy` accepts a description and a typed predicate:

```go
c.Expect(4).To(Satisfy("be even", func(value int) bool {
	return value%2 == 0
}))
```

See the executable [expectations example](./examples/expect_test.go) for every built-in matcher.

## Accessing `testing.T`

`Case.T` returns the underlying `*testing.T` when a test needs an API outside gspec:

```go
c.It("uses a testing helper", func(c *gspec.Case) {
	somethingThatNeedsTestingT(c.T())
})
```

## RSpec feature comparison

| Feature | `gspec` |
| --- | --- |
| Example groups | Supported with `Describe`, `Context`, and `It` |
| Lazy values | Supported with `Let`, `Get`, and `Set` |
| Hooks | Supported with `BeforeEach` and `AfterEach` |
| Fluent expectations | Supported with `Expect`, `To`, and `NotTo` |
| Matchers | Built-in typed matchers and custom predicates with `Satisfy` |
| Mocks | Use a dedicated Go mocking library |

## Why?

Go's built-in testing tools are intentionally small and effective. The value of gspec is organization: related cases can share descriptions, hooks, and declarations without sharing mutable values.

Each `Let` is lazy, cached only within its current case, and overridable in a nested group. This makes it possible to describe variations of a subject without manually coordinating closure state between subtests.

More complete, runnable demonstrations are available in the [`examples` directory](./examples).
