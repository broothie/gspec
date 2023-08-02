package gspec

import (
	"testing"

	"github.com/broothie/gspec/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLet(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		mockT := mocks.NewMocktestingT(gomock.NewController(t))
		mockT.EXPECT().Helper().AnyTimes()

		allowTestFuncs(mockT, "behavior")

		Run(mockT, func(c *Context) {
			something := Let(c, "something", func(c *Case) string { return "first" })

			c.It("behavior", func(c *Case) {
				assert.Equal(t, "first", something(c))
			})
		})
	})

	t.Run("parent", func(t *testing.T) {
		mockT := mocks.NewMocktestingT(gomock.NewController(t))
		mockT.EXPECT().Helper().AnyTimes()

		allowTestFuncs(mockT, "behavior")
		allowTestFuncs(mockT, "nested behavior")

		Run(mockT, func(c *Context) {
			something := Let(c, "something", func(c *Case) string { return "first" })

			c.It("behavior", func(c *Case) {
				assert.Equal(t, "first", something(c))
			})

			c.Describe("nested", func(c *Context) {
				c.It("behavior", func(c *Case) {
					assert.Equal(t, "first", something(c))
				})
			})
		})
	})

	t.Run("nested", func(t *testing.T) {
		mockT := mocks.NewMocktestingT(gomock.NewController(t))
		mockT.EXPECT().Helper().AnyTimes()

		allowTestFuncs(mockT, "behavior")
		allowTestFuncs(mockT, "nested behavior")

		Run(mockT, func(c *Context) {
			something := Let(c, "something", func(c *Case) string { return "first" })

			c.It("behavior", func(c *Case) {
				assert.Equal(t, "first", something(c))
			})

			c.Describe("nested", func(c *Context) {
				something := Let(c, "something", func(c *Case) string { return "second" })

				c.It("behavior", func(c *Case) {
					assert.Equal(t, "second", something(c))
				})
			})
		})
	})

	t.Run("nested without function shadowing", func(t *testing.T) {
		mockT := mocks.NewMocktestingT(gomock.NewController(t))
		mockT.EXPECT().Helper().AnyTimes()

		allowTestFuncs(mockT, "behavior")
		allowTestFuncs(mockT, "nested behavior")

		Run(mockT, func(c *Context) {
			something := Let(c, "something", func(c *Case) string { return "first" })

			c.It("behavior", func(c *Case) {
				assert.Equal(t, "first", something(c))
			})

			c.Describe("nested", func(c *Context) {
				Let(c, "something", func(c *Case) string { return "second" })

				c.It("behavior", func(c *Case) {
					assert.Equal(t, "second", something(c))
				})
			})
		})
	})

	t.Run("caching", func(t *testing.T) {
		mockT := mocks.NewMocktestingT(gomock.NewController(t))
		mockT.EXPECT().Helper().AnyTimes()

		allowTestFuncs(mockT, "behavior")

		calls := 0
		Run(mockT, func(c *Context) {
			something := Let(c, "something", func(c *Case) string {
				assert.Equal(t, calls, 0)
				calls++
				return "first"
			})

			c.It("behavior", func(c *Case) {
				assert.Equal(t, "first", something(c))
				assert.Equal(t, "first", something(c))
			})
		})

		assert.Equal(t, calls, 1)
	})
}

func TestContext_findLet(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		context := &Context{
			lets: map[string]letFunc{
				"some-let": func(c *Case) any { return "child value" },
			},
			parent: &Context{
				lets: map[string]letFunc{
					"some-let": func(c *Case) any { return "parent value" },
				},
			},
		}

		assert.Equal(t, "child value", context.findLet("some-let")(nil))
	})

	t.Run("found in parent", func(t *testing.T) {
		context := &Context{
			lets: make(map[string]letFunc),
			parent: &Context{
				lets: map[string]letFunc{
					"some-let": func(c *Case) any { return "parent value" },
				},
			},
		}

		assert.Equal(t, "parent value", context.findLet("some-let")(nil))
	})

	t.Run("undefined", func(t *testing.T) {
		context := &Context{
			lets:   make(map[string]letFunc),
			parent: &Context{lets: make(map[string]letFunc)},
		}

		assert.Panics(t, func() { context.findLet("some-let") })
	})
}
