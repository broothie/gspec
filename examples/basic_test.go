package examples

import (
	"testing"

	"github.com/broothie/gspec"
	"github.com/broothie/gspec/testhelp"
)

func Test(t *testing.T) {
	gspec.Describe(t, "addition", func(c *gspec.Context) {
		c.It("returns the sum of its operands", func(c *gspec.Case) {
			testhelp.AssertEqual(c.T(), 3, 1+2)
		})
	})
}
