package examples

import (
	"testing"

	"github.com/broothie/gspec"
	. "github.com/broothie/gspec/match"
)

func TestGroups(t *testing.T) {
	gspec.Run(t, func(c *gspec.Context) {
		c.Describe("strings", func(c *gspec.Context) {
			c.Context("when text is present", func(c *gspec.Context) {
				c.It("can describe its behavior", func(c *gspec.Case) {
					c.Expect("hello").NotTo(Equal(""))
				})
			})
		})
	})
}
