package gspec

import (
	"testing"

	"github.com/broothie/gspec/testhelp"
)

func TestTestContext(t *testing.T) {
	mockT := testhelp.NewTestingTMock(t)
	mockT.ExpectRun("subject context behavior")

	Describe(mockT, "subject", func(c *TestContext) {
		c.Context("context", func(c *TestContext) {
			c.It("behavior", func(c *TestCase) {})
		})
	})
}

func TestContext_joinNames(t *testing.T) {
	context := &TestContext{
		name: "bottom",
		parent: &TestContext{
			name: "middle",
			parent: &TestContext{
				name: "top",
			},
		},
	}

	testhelp.AssertEqual(t, "top middle bottom case", context.joinNames("case"))
}
