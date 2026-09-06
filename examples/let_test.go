package examples

import (
	"strings"
	"testing"

	"github.com/broothie/gspec"
	"github.com/broothie/gspec/testhelp"
)

func capitalize(input string) string {
	return strings.ToUpper(input)
}

func Test_capitalize(t *testing.T) {
	gspec.Run(t, func(c *gspec.Context) {
		input := c.Let(func(c *gspec.Case) string { return "Hello" })

		c.It("should capitalize the input", func(c *gspec.Case) {
			testhelp.AssertEqual(c.T(), "HELLO", capitalize(c.Get(input)))
		})

		c.Context("with spaces", func(c *gspec.Context) {
			input := c.Let(func(c *gspec.Case) string { return "Hello, world" })

			c.It("should capitalize the input", func(c *gspec.Case) {
				testhelp.AssertEqual(c.T(), "HELLO, WORLD", capitalize(c.Get(input)))
			})
		})
	})
}
