package gspec

import (
	"testing"

	"github.com/broothie/gspec/testhelp"
)

func TestContext(t *testing.T) {
	mockT := testhelp.NewTestingTMock(t)
	mockT.ExpectRun("subject context behavior")

	Describe(mockT, "subject", func(c *Context) {
		c.Context("context", func(c *Context) {
			c.It("behavior", func(c *Case) {})
		})
	})
}

func TestContext_joinNames(t *testing.T) {
	context := &Context{
		name: "bottom",
		parent: &Context{
			name: "middle",
			parent: &Context{
				name: "top",
			},
		},
	}

	testhelp.AssertEqual(t, "top middle bottom case", context.joinNames("case"))
}

func Test_joinNames(t *testing.T) {
	t.Run("receiver first", func(t *testing.T) {
		strs := []string{"*Object", ".method", "when some context", "behaves some way"}

		testhelp.AssertEqual(t, "*Object.method when some context behaves some way", joinNames(strs...))
	})

	t.Run("receiver in the middle", func(t *testing.T) {
		strs := []string{"objects", "*Object", ".method", "behaves some way"}

		testhelp.AssertEqual(t, "objects *Object.method behaves some way", joinNames(strs...))
	})

	t.Run("empty value in the middle", func(t *testing.T) {
		strs := []string{"objects", "*Object", ".method", "", "behaves some way"}

		testhelp.AssertEqual(t, "objects *Object.method behaves some way", joinNames(strs...))
	})
}
