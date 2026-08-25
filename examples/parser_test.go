package examples

import (
	"testing"

	"github.com/broothie/gspec"
	"github.com/broothie/gspec/testhelp"
)

type Parser struct {
	index  int
	tokens []string
}

func (p *Parser) IsExhausted() bool {
	return p.index >= len(p.tokens)
}

func Test_advanced_let(t *testing.T) {
	gspec.Describe(t, "Parser", func(c *gspec.Context) {
		tokens := c.Let(func(c *gspec.Case) []string {
			return []string{"arg1", "arg2", "-f", "filename"}
		})

		parser := c.Let(func(c *gspec.Case) *Parser { return &Parser{tokens: c.Get(tokens)} })

		c.Describe(".IsExhausted", func(c *gspec.Context) {
			c.Context("when tokens remain", func(c *gspec.Context) {
				c.It("is false", func(c *gspec.Case) {
					testhelp.AssertEqual(c.T(), c.Get(parser).IsExhausted(), false)
				})
			})

			c.Context("when no tokens remain", func(c *gspec.Context) {
				c.BeforeEach(func(c *gspec.Case) {
					c.Get(parser).index = 4
				})

				c.It("is true", func(c *gspec.Case) {
					testhelp.AssertEqual(c.T(), c.Get(parser).IsExhausted(), true)
				})
			})

			c.Context("when tokens is empty", func(c *gspec.Context) {
				c.Set(tokens, func(c *gspec.Case) []string { return nil })

				c.It("is true", func(c *gspec.Case) {
					testhelp.AssertEqual(c.T(), c.Get(parser).IsExhausted(), true)
				})
			})
		})
	})
}
