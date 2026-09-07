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
	gspec.Run(t, func(t *gspec.TestContext) {
		mux := t.Let(func(t *gspec.TestCase) *http.ServeMux { return http.NewServeMux() })
		server := t.Let(func(t *gspec.TestCase) *httptest.Server { return httptest.NewServer(t.Get(mux)) })

		t.BeforeEach(func(t *gspec.TestCase) {
			t.Get(mux).HandleFunc("/api/teapot", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusTeapot)
			})
		})

		t.AfterEach(func(t *gspec.TestCase) {
			t.Get(server).Close()
		})

		t.It("serves requests", func(t *gspec.TestCase) {
			response, err := http.Get(fmt.Sprintf("%s/api/teapot", t.Get(server).URL))
			t.Expect(err).NotTo(HaveOccurred())
			if err != nil {
				return
			}
			defer response.Body.Close()

			t.Expect(response.StatusCode).To(Equal(http.StatusTeapot))
		})
	})
}
