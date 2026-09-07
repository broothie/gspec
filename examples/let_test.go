package examples

import (
	"strings"
	"testing"

	"github.com/broothie/gspec"
	. "github.com/broothie/gspec/match"
)

func capitalize(input string) string {
	return strings.ToUpper(input)
}

func TestLet(t *testing.T) {
	gspec.Run(t, func(c *gspec.Context) {
		input := c.Let(func(c *gspec.Case) string { return "Hello" })

		c.It("evaluates a value for the test case", func(c *gspec.Case) {
			c.Expect(capitalize(c.Get(input))).To(Equal("HELLO"))
		})

		c.Context("with spaces", func(c *gspec.Context) {
			c.Set(input, func(c *gspec.Case) string { return "Hello, world" })

			c.It("uses the value overridden by the context", func(c *gspec.Case) {
				c.Expect(capitalize(c.Get(input))).To(Equal("HELLO, WORLD"))
			})
		})
	})
}
