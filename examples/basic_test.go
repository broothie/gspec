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
