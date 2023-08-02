package gspec

import (
	"testing"

	"github.com/broothie/gspec/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestContext(t *testing.T) {
	mockT := mocks.NewMocktestingT(gomock.NewController(t))
	mockT.EXPECT().Helper().AnyTimes()

	allowTestFuncs(mockT, "subject context behavior")

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

	assert.Equal(t, "top middle bottom case", context.joinNames("case"))
}
