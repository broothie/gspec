package examples

import (
	"strings"
	"testing"

	"github.com/broothie/gspec"
)

func capitalize(input string) string {
	return strings.ToUpper(input)
}

func Test_capitalize(t *testing.T) {
	gspec.Run(t, func(c *gspec.Context) {
		input := gspec.Let(c, "input", func(c *gspec.Case) string { return "Hello" })

		c.It("should capitalize the input", func(c *gspec.Case) {
			c.Assert().Equal("HELLO", capitalize(input(c)))
		})

		c.Context("with spaces", func(c *gspec.Context) {
			gspec.Let(c, "input", func(c *gspec.Case) string { return "Hello, world" })

			c.It("should capitalize the input", func(c *gspec.Case) {
				c.Assert().Equal("HELLO, WORLD", capitalize(input(c)))
			})
		})
	})
}
