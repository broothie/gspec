package gspec

import (
	"testing"

	"github.com/broothie/gspec/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestContext_runCase(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		mockT := mocks.NewMocktestingT(gomock.NewController(t))
		mockT.EXPECT().Helper().Times(1)

		called := false

		c := &Context{name: "some context"}
		c.runCase(mockT, caseEntry{
			name: "case",
			run: func(c *Case) {
				called = true
			},
		})

		assert.True(t, called)
	})

	t.Run("with hooks", func(t *testing.T) {
		mockT := mocks.NewMocktestingT(gomock.NewController(t))
		mockT.EXPECT().Helper().Times(1)

		calls := 0

		c := &Context{
			name: "some context",
			befores: []CaseFunc{func(c *Case) {
				assert.Equal(t, calls, 0)
				calls++
			}},
			afters: []CaseFunc{func(c *Case) {
				assert.Equal(t, calls, 2)
				calls++
			}},
		}

		c.runCase(mockT, caseEntry{
			name: "case",
			run: func(c *Case) {
				assert.Equal(t, calls, 1)
				calls++
			},
		})

		assert.Equal(t, calls, 3)
	})
}
