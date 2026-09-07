package examples

import (
	"testing"

	"github.com/broothie/gspec"
	. "github.com/broothie/gspec/match"
)

type Parser struct {
	index  int
	tokens []string
}

func (p *Parser) IsExhausted() bool {
	return p.index >= len(p.tokens)
}

func TestAdvancedLet(t *testing.T) {
	gspec.Describe(t, "Parser", func(t *gspec.TestContext) {
		tokens := t.Let(func(t *gspec.TestCase) []string {
			return []string{"arg1", "arg2", "-f", "filename"}
		})

		parser := t.Let(func(t *gspec.TestCase) *Parser { return &Parser{tokens: t.Get(tokens)} })

		t.Describe(".IsExhausted", func(t *gspec.TestContext) {
			t.Context("when tokens remain", func(t *gspec.TestContext) {
				t.It("is false", func(t *gspec.TestCase) {
					t.Expect(t.Get(parser).IsExhausted()).To(Equal(false))
				})
			})

			t.Context("when no tokens remain", func(t *gspec.TestContext) {
				t.BeforeEach(func(t *gspec.TestCase) {
					t.Get(parser).index = 4
				})

				t.It("is true", func(t *gspec.TestCase) {
					t.Expect(t.Get(parser).IsExhausted()).To(Equal(true))
				})
			})

			t.Context("when tokens is empty", func(t *gspec.TestContext) {
				t.Set(tokens, func(t *gspec.TestCase) []string { return nil })

				t.It("is true", func(t *gspec.TestCase) {
					t.Expect(t.Get(parser).IsExhausted()).To(Equal(true))
				})
			})
		})
	})
}
