package gspec

import (
	"testing"

	"github.com/broothie/gspec/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestContext_BeforeEach(t *testing.T) {
	mockT := mocks.NewMocktestingT(gomock.NewController(t))
	mockT.EXPECT().Helper().AnyTimes()

	allowTestFuncs(mockT, "tests")

	calls := 0
	Run(mockT, func(c *Context) {
		c.BeforeEach(func(c *Case) {
			assert.Equal(t, calls, 0)
			calls++
		})

		c.It("tests", func(c *Case) {
			assert.Equal(t, calls, 1)
			calls++
		})
	})

	assert.Equal(t, calls, 2)
}

func TestContext_AfterEach(t *testing.T) {
	mockT := mocks.NewMocktestingT(gomock.NewController(t))
	mockT.EXPECT().Helper().AnyTimes()

	allowTestFuncs(mockT, "tests")

	calls := 0
	Run(mockT, func(c *Context) {
		c.AfterEach(func(c *Case) {
			assert.Equal(t, calls, 1)
			calls++
		})

		c.It("tests", func(c *Case) {
			assert.Equal(t, calls, 0)
			calls++
		})
	})

	assert.Equal(t, calls, 2)
}
