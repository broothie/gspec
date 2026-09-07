package examples

import (
	"testing"

	"github.com/broothie/gspec"
	. "github.com/broothie/gspec/match"
)

func TestBasic(t *testing.T) {
	gspec.Describe(t, "addition", func(t *gspec.TestContext) {
		t.It("returns the sum of its operands", func(t *gspec.TestCase) {
			t.Expect(1 + 2).To(Equal(3))
		})
	})
}
