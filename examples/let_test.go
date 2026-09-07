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
	gspec.Run(t, func(t *gspec.TestContext) {
		input := t.Let(func(t *gspec.TestCase) string { return "Hello" })

		t.It("evaluates a value for the test case", func(t *gspec.TestCase) {
			t.Expect(capitalize(t.Get(input))).To(Equal("HELLO"))
		})

		t.Context("with spaces", func(t *gspec.TestContext) {
			t.Set(input, func(t *gspec.TestCase) string { return "Hello, world" })

			t.It("uses the value overridden by the context", func(t *gspec.TestCase) {
				t.Expect(capitalize(t.Get(input))).To(Equal("HELLO, WORLD"))
			})
		})
	})
}
