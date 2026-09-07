package examples

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/broothie/gspec"
	. "github.com/broothie/gspec/match"
)

func TestHooks(t *testing.T) {
	gspec.Run(t, func(c *gspec.Context) {
		mux := c.Let(func(c *gspec.Case) *http.ServeMux { return http.NewServeMux() })
		server := c.Let(func(c *gspec.Case) *httptest.Server { return httptest.NewServer(c.Get(mux)) })

		c.BeforeEach(func(c *gspec.Case) {
			c.Get(mux).HandleFunc("/api/teapot", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusTeapot)
			})
		})

		c.AfterEach(func(c *gspec.Case) {
			c.Get(server).Close()
		})

		c.It("serves requests", func(c *gspec.Case) {
			response, err := http.Get(fmt.Sprintf("%s/api/teapot", c.Get(server).URL))
			c.Expect(err).NotTo(HaveOccurred())
			if err != nil {
				return
			}
			defer response.Body.Close()

			c.Expect(response.StatusCode).To(Equal(http.StatusTeapot))
		})
	})
}
