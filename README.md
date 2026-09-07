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

```go
package calculator_test

import (
	"testing"

	"github.com/broothie/gspec"
	. "github.com/broothie/gspec/match" // Matchers are designed to read naturally when the `match` package is dot-imported
)

func TestAddition(t *testing.T) {
	gspec.Describe(t, "addition", func(t *gspec.TestContext) {
		t.It("returns the sum of its operands", func(t *gspec.TestCase) {
			t.Expect(1 + 2).To(Equal(3))
			t.Expect(1 + 2).NotTo(Equal(4))
		})
	})
}
```

`gspec.Describe` opens a labelled root group. Use `gspec.Run` when the root does not need a label. Each call to `It` becomes a normal Go subtest, so existing `go test` tooling continues to work.

## Groups

Use `Describe` to name a subject and `Context` to describe a condition. Groups can be nested arbitrarily and inherit lets and hooks from their parents.

```go
func TestGreeting(t *testing.T) {
	gspec.Run(t, func(t *gspec.TestContext) {
		t.Describe("Greeting", func(t *gspec.TestContext) {
			t.Context("when a name is present", func(t *gspec.TestContext) {
				t.It("includes the name", func(t *gspec.TestCase) {
					t.Expect("Hello, Gopher!").To(ContainSubstring("Gopher"))
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
	gspec.Describe(t, "Greeting", func(t *gspec.TestContext) {
		name := t.Let(func(t *gspec.TestCase) string { return "Gopher" })
		greeting := t.Let(func(t *gspec.TestCase) string {
			return "Hello, " + t.Get(name) + "!"
		})

		t.It("greets the default name", func(t *gspec.TestCase) {
			t.Expect(t.Get(greeting)).To(Equal("Hello, Gopher!"))
		})

		t.Context("with another name", func(t *gspec.TestContext) {
			t.Set(name, func(t *gspec.TestCase) string { return "Rubyist" })

			t.It("uses the overridden name", func(t *gspec.TestCase) {
				t.Expect(t.Get(greeting)).To(Equal("Hello, Rubyist!"))
			})
		})
	})
}
```

## Hooks

`BeforeEach` and `AfterEach` register setup and cleanup functions. Hooks are inherited by nested groups and receive the current `TestCase`, allowing them to access lets with `Get`.

```go
gspec.Run(t, func(t *gspec.TestContext) {
	values := t.Let(func(t *gspec.TestCase) *[]int { return &[]int{} })

	t.BeforeEach(func(t *gspec.TestCase) {
		caseValues := t.Get(values)
		*caseValues = append(*caseValues, 1)
	})

	t.AfterEach(func(t *gspec.TestCase) {
		*t.Get(values) = nil
	})

	t.It("runs between the hooks", func(t *gspec.TestCase) {
		t.Expect(*t.Get(values)).To(Contain(1))
	})
})
```

See the executable [hooks example](./examples/hooks_test.go) for a complete setup and cleanup flow.

## Expectations

Inside a case, `Expect` captures an actual value. `To` reports a failure when its matcher does not match, while `NotTo` reports a failure when it does.

```go
t.Expect(actual).To(Equal(expected))
t.Expect(actual).NotTo(Equal(unexpected))
```

Outside a `TestCase`, the package-level form reports to any compatible test value:

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
t.Expect(pointer).To(BeNil[*Widget]())
t.Expect(values).To(BeEmpty[[]int]())
t.Expect(values).To(HaveLength[[]int](3))
t.Expect(err).To(BeErrorType[*ValidationError]())
```

For a one-off assertion, `Satisfy` accepts a description and a typed predicate:

```go
t.Expect(4).To(Satisfy("be even", func(value int) bool {
	return value%2 == 0
}))
```

See the executable [expectations example](./examples/expect_test.go) for every built-in matcher.

## Using `testing.T`

`TestCase` embeds its underlying `*testing.T`, so standard testing methods are available directly. The embedded `T` field can be passed to functions that specifically require a `*testing.T`:

```go
t.It("uses a testing helper", func(t *gspec.TestCase) {
	t.Cleanup(cleanup)
	somethingThatNeedsTestingT(t.T)
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
