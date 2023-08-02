package examples

import (
	"testing"

	"github.com/broothie/gspec"
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
		tokens := gspec.Let(c, "tokens", func(c *gspec.Case) []string {
			return []string{"arg1", "arg2", "-f", "filename"}
		})

		parser := gspec.Let(c, "parser", func(c *gspec.Case) *Parser { return &Parser{tokens: tokens(c)} })

		c.Describe(".IsExhausted", func(c *gspec.Context) {
			c.Context("when tokens remain", func(c *gspec.Context) {
				c.It("is false", func(c *gspec.Case) {
					c.Assert().False(parser(c).IsExhausted())
				})
			})

			c.Context("when no tokens remain", func(c *gspec.Context) {
				c.BeforeEach(func(c *gspec.Case) {
					parser(c).index = 4
				})

				c.It("is true", func(c *gspec.Case) {
					c.Assert().True(parser(c).IsExhausted())
				})
			})

			c.Context("when tokens is empty", func(c *gspec.Context) {
				gspec.Let(c, "tokens", func(c *gspec.Case) []string { return nil })

				c.It("is true", func(c *gspec.Case) {
					c.Assert().True(parser(c).IsExhausted())
				})
			})
		})
	})
}
