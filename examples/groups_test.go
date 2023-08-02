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
