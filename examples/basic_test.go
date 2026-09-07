package examples

import (
	"testing"

	"github.com/broothie/gspec"
	. "github.com/broothie/gspec/match"
)

func TestBasic(t *testing.T) {
	gspec.Describe(t, "addition", func(c *gspec.Context) {
		c.It("returns the sum of its operands", func(c *gspec.Case) {
			c.Expect(1 + 2).To(Equal(3))
		})
	})
}
