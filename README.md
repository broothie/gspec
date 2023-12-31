# gspec

[![Go project version](https://badge.fury.io/go/github.com%2Fbroothie%2Fgspec.svg)](https://badge.fury.io/go/github.com%2Fbroothie%2Fgspec)
[![Go Report Card](https://goreportcard.com/badge/github.com/broothie/gspec)](https://goreportcard.com/report/github.com/broothie/gspec)
[![codecov](https://codecov.io/gh/broothie/gspec/branch/main/graph/badge.svg?token=6CLN4sDTk5)](https://codecov.io/gh/broothie/gspec)
[![gosec](https://github.com/broothie/gspec/actions/workflows/gosec.yml/badge.svg)](https://github.com/broothie/gspec/actions/workflows/gosec.yml)
[![GitHub](https://img.shields.io/github/license/broothie/gspec)](https://opensource.org/license/mit/)

`gspec` is a testing framework for Go, inspired by Ruby's [`rspec`](http://rspec.info).

## Installation

```shell
go get github.com/broothie/gspec
```

## Usage

### Basics

`gspec` hooks into Go's built-in testing framework.
1. In regular Go test function, `gspec.Describe` or `gspec.Run` are used to open a `gspec` context.
2. Then, `c.It` is used to define an actual test case.
3. Within a test case, `c.Assert()` returns an
   [`*assert.Assertions`](https://pkg.go.dev/github.com/stretchr/testify@v1.8.4/assert#Assertions),
   which can be used to make assertions about the code under test.

```go
package examples

import (
   "testing"

   "github.com/broothie/gspec"
)

func Test(t *testing.T) {
   gspec.Describe(t, "addition", func(c *gspec.Context) {
      c.It("returns the sum of its operands", func(c *gspec.Case) {
         c.Assert().Equal(3, 1+2)
      })
   })
}
```

If you need to access the underlying `*testing.T`, you can do so from within a hook or test case via `c.T()`.

```go
package examples

import (
   "testing"

   "github.com/broothie/gspec"
)

func somethingThatNeedsTestingT(t *testing.T) {}

func Test_t(t *testing.T) {
   gspec.Describe(t, ".T", func(c *gspec.Context) {
      c.It("returns a *testing.T", func(c *gspec.Case) {
         somethingThatNeedsTestingT(c.T())
      })
   })
}
```

### Groups

Test cases can be grouped together via `c.Describe` and `c.Context`.
Groups can be nested arbitrarily.
Groups inherit [`Let`](#let)s and [hooks](#hooks) from their parents.

```go
package examples

import (
   "testing"

   "github.com/broothie/gspec"
)

func Test_groups(t *testing.T) {
   gspec.Run(t, func(c *gspec.Context) {
      c.Describe("some subject", func(c *gspec.Context) {
         c.Context("when in some context", func(c *gspec.Context) {
            c.It("does something", func(c *gspec.Case) {
               // Test code, assertions, etc.
            })
         })
      })
   })
}
```

### Let

`gspec.Let` allows for the definition of type-safe, per-case values.
`Let` values are only evaluated if they are used in a test case,
and are cached for the duration of the test case.

`Let` values can be overwritten in nested groups, but **their return type must remain the same**.
When overwriting a `Let` in this way, the returned function needn't be captured.
The value will still be registered for the context, even though the function was captured in an outer group.

```go
package examples

import (
   "strings"
   "testing"

   "github.com/broothie/gspec"
)

func capitalize(input string) string {
   return strings.ToUpper(input)
}

func Test_capitalize(t *testing.T) {
   gspec.Run(t, func(c *gspec.Context) {
      input := gspec.Let(c, "input", func(c *gspec.Case) string { return "Hello" })

      c.It("should capitalize the input", func(c *gspec.Case) {
         c.Assert().Equal("HELLO", capitalize(input(c)))
      })

      c.Context("with spaces", func(c *gspec.Context) {
         gspec.Let(c, "input", func(c *gspec.Case) string { return "Hello, world" })

         c.It("should capitalize the input", func(c *gspec.Case) {
            c.Assert().Equal("HELLO, WORLD", capitalize(input(c)))
         })
      })
   })
}
```

### Hooks

`c.BeforeEach` and `c.AfterEach` can be used to register hooks that run around each test case.


Hooks are inherited by nested groups.

```go
package examples

import (
   "fmt"
   "net/http"
   "net/http/httptest"
   "testing"

   "github.com/broothie/gspec"
)

func Test_hooks(t *testing.T) {
   gspec.Run(t, func(c *gspec.Context) {
      mux := gspec.Let(c, "mux", func(c *gspec.Case) *http.ServeMux { return http.NewServeMux() })
      server := gspec.Let(c, "server", func(c *gspec.Case) *httptest.Server { return httptest.NewServer(mux(c)) })

      c.BeforeEach(func(c *gspec.Case) {
         mux(c).HandleFunc("/api/teapot", func(w http.ResponseWriter, r *http.Request) {
            w.WriteHeader(http.StatusTeapot)
         })
      })

      c.AfterEach(func(c *gspec.Case) {
         server(c).Close()
      })

      c.It("serves requests", func(c *gspec.Case) {
         response, err := http.Get(fmt.Sprintf("%s/api/teapot", server(c).URL))
         c.Assert().NoError(err)
         c.Assert().Equal(http.StatusTeapot, response.StatusCode)
      })
   })
}
```

## RSpec Feature Comparison

| Feature                    | `gspec`                                                                                                                |
|----------------------------|------------------------------------------------------------------------------------------------------------------------|
| Example Groups             | ✅                                                                                                                      |
| Let                        | ✅                                                                                                                      |
| Hooks                      | ✅                                                                                                                      |
| Mocks                      | Use an existing mock library, such as https://github.com/uber-go/mock.                                                 |
| Fluent-syntax expectations | `*gspec.Case` exposes assertions from [`assert`](https://github.com/stretchr/testify#assert-package) via `c.Assert()`. |

## Why?

Go's built-in testing utilities are pretty good on their own.
Paired with a library like [`assert`](https://github.com/stretchr/testify#assert-package) and Go testing is pretty dang good.

I think the power of this package comes from [`Let`](#let), and how it works with [groups](#groups).
Go's `t.Run` and its use of closures makes it difficult/confusing to define reusable values in an outer scope which can be overwritten in an inner scope.
Plus, having multiple tests that close over the same value runs the risk of modification of that shared value.

`Let` values are per-case, lazy-evaluated, overwrite-able, and cached for the duration of the test case.
Since they're overwrite-able, a `Let` can be redefined for a subgroup, even if they're not specifically referenced from within that group's test cases.

```go
package examples

import (
  "testing"

  "github.com/broothie/gspec"
)

type Parser struct {
  index  int
  tokens []string
}

func (p *Parser) IsExhausted() bool {
  return p.index >= len(p.tokens)
}

func Test_advanced_let(t *testing.T) {
  gspec.Describe(t, "Parser", func(c *gspec.Context) {
    tokens := gspec.Let(c, "tokens", func(c *gspec.Case) []string {
      return []string{"arg1", "arg2", "-f", "filename"}
    })

    parser := gspec.Let(c, "parser", func(c *gspec.Case) *Parser { return &Parser{tokens: tokens(c)} })

    c.Describe(".IsExhausted", func(c *gspec.Context) {
      c.Context("when tokens remain", func(c *gspec.Context) {
        c.It("is false", func(c *gspec.Case) {
          c.Assert().False(parser(c).IsExhausted())
        })
      })

      c.Context("when no tokens remain", func(c *gspec.Context) {
        c.BeforeEach(func(c *gspec.Case) {
          parser(c).index = 4
        })

        c.It("is true", func(c *gspec.Case) {
          c.Assert().True(parser(c).IsExhausted())
        })
      })

      c.Context("when tokens is empty", func(c *gspec.Context) {
        gspec.Let(c, "tokens", func(c *gspec.Case) []string { return nil })

        c.It("is true", func(c *gspec.Case) {
          c.Assert().True(parser(c).IsExhausted())
        })
      })
    })
  })
}
```
