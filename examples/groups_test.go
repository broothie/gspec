package examples

import (
	"testing"

	"github.com/broothie/gspec"
	. "github.com/broothie/gspec/match"
)

func TestGroups(t *testing.T) {
	gspec.Run(t, func(t *gspec.TestContext) {
		t.Describe("strings", func(t *gspec.TestContext) {
			t.Context("when text is present", func(t *gspec.TestContext) {
				t.It("can describe its behavior", func(t *gspec.TestCase) {
					t.Expect("hello").NotTo(Equal(""))
				})
			})
		})
	})
}
