package gspec

import (
	"testing"

	"github.com/broothie/gspec/testhelp"
)

func TestLet(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		mockT := testhelp.NewTestingTMock(t)
		mockT.ExpectRun("behavior")

		Run(mockT, func(c *Context) {
			something := c.Let(func(c *Case) string { return "first" })

			c.It("behavior", func(c *Case) {
				testhelp.AssertEqual(t, "first", c.Get(something))
			})
		})
	})

	t.Run("parent", func(t *testing.T) {
		mockT := testhelp.NewTestingTMock(t)
		mockT.ExpectRun("behavior")
		mockT.ExpectRun("nested behavior")

		Run(mockT, func(c *Context) {
			something := c.Let(func(c *Case) string { return "first" })

			c.It("behavior", func(c *Case) {
				testhelp.AssertEqual(t, "first", c.Get(something))
			})

			c.Describe("nested", func(c *Context) {
				c.It("behavior", func(c *Case) {
					testhelp.AssertEqual(t, "first", c.Get(something))
				})
			})
		})
	})

	t.Run("nested", func(t *testing.T) {
		mockT := testhelp.NewTestingTMock(t)
		mockT.ExpectRun("behavior")
		mockT.ExpectRun("nested behavior")

		Run(mockT, func(c *Context) {
			something := c.Let(func(c *Case) string { return "first" })

			c.It("behavior", func(c *Case) {
				testhelp.AssertEqual(t, "first", c.Get(something))
			})

			c.Describe("nested", func(c *Context) {
				something := c.Let(func(c *Case) string { return "second" })

				c.It("behavior", func(c *Case) {
					testhelp.AssertEqual(t, "second", c.Get(something))
				})
			})
		})
	})

	t.Run("caching", func(t *testing.T) {
		mockT := testhelp.NewTestingTMock(t)
		mockT.ExpectRun("behavior")

		calls := 0
		Run(mockT, func(c *Context) {
			something := c.Let(func(c *Case) string {
				testhelp.AssertEqual(t, calls, 0)
				calls++
				return "first"
			})

			c.It("behavior", func(c *Case) {
				testhelp.AssertEqual(t, "first", c.Get(something))
				testhelp.AssertEqual(t, "first", c.Get(something))
			})
		})

		testhelp.AssertEqual(t, calls, 1)
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

		testhelp.AssertEqual(t, "child value", context.findLet("some-let")(nil))
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

		testhelp.AssertEqual(t, "parent value", context.findLet("some-let")(nil))
	})

	t.Run("undefined", func(t *testing.T) {
		context := &Context{
			lets:   make(map[string]letFunc),
			parent: &Context{lets: make(map[string]letFunc)},
		}

		recoverCalled := false
		defer func() {
			testhelp.AssertEqual(t, recoverCalled, true)
		}()

		defer func() {
			if value := recover(); value != nil {
				recoverCalled = true
			}
		}()

		context.findLet("some-let")
	})
}
