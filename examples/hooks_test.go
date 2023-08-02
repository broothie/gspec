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
