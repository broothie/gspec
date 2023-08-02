# gspec

`gspec` is a testing framework for Go, inspired by Ruby's [`rspec`](http://rspec.info).

## Installation

```shell
go get github.com/broothie/gspec
```

## Usage

### Basics

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
	})
}
```

### Groups

```go
```

### Let

`gspec.Let` allows you to define type-safe, per-case values.
Let values are cached for the duration of the test case.

```go

```

### Hooks

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
		client := gspec.Let(c, "client", func(c *gspec.Case) *http.Client { return server(c).Client() })
		url := gspec.Let(c, "url", func(c *gspec.Case) string { return server(c).URL })

		c.BeforeEach(func(c *gspec.Case) {
			mux(c).HandleFunc("/api/teapot", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusTeapot)
			})
		})

		c.AfterEach(func(c *gspec.Case) {
			server(c).Close()
		})

		c.It("serves requests", func(c *gspec.Case) {
			response, err := client(c).Get(fmt.Sprintf("%s/api/teapot", url(c)))
			c.Assert().NoError(err)
			c.Assert().Equal(http.StatusTeapot, response.StatusCode)
		})
	})
}
```

## RSpec Feature Parity

| Feature                    | `rspec` | `gspec`                                                                                                        |
|----------------------------|---------|----------------------------------------------------------------------------------------------------------------|
| Example Groups             | ✅       | ✅                                                                                                              |
| Let                        | ✅       | ✅                                                                                                              |
| Hooks                      | ✅       | ✅                                                                                                              |
| Mocks                      | ✅       | Too difficult to build in. Use an existing mock library, such as https://github.com/uber-go/mock.              |
| Fluent-syntax expectations | ✅       | `*gspec.Case` exposes assertions from [stretchr/testify](https://github.com/stretchr/testify) via `.Assert()`. |


## Why?

Go's `*testing.T` does a lot on its own.
Paired with a package like [`testify`](https://github.com/stretchr/testify) and you've got pretty much all you need for writing basic tests.

So why would we want a framework like `gspec`?


