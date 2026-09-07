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
