package gspec

import (
	"testing"

	"github.com/broothie/gspec/testhelp"
)

func TestLet(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		mockT := testhelp.NewTestingTMock(t)
		mockT.ExpectRun("behavior")

		Run(mockT, func(c *TestContext) {
			something := c.Let(func(c *TestCase) string { return "first" })

			c.It("behavior", func(c *TestCase) {
				testhelp.AssertEqual(t, "first", c.Get(something))
			})
		})
	})

	t.Run("parent", func(t *testing.T) {
		mockT := testhelp.NewTestingTMock(t)
		mockT.ExpectRun("behavior")
		mockT.ExpectRun("nested behavior")

		Run(mockT, func(c *TestContext) {
			something := c.Let(func(c *TestCase) string { return "first" })

			c.It("behavior", func(c *TestCase) {
				testhelp.AssertEqual(t, "first", c.Get(something))
			})

			c.Describe("nested", func(c *TestContext) {
				c.It("behavior", func(c *TestCase) {
					testhelp.AssertEqual(t, "first", c.Get(something))
				})
			})
		})
	})

	t.Run("nested", func(t *testing.T) {
		mockT := testhelp.NewTestingTMock(t)
		mockT.ExpectRun("behavior")
		mockT.ExpectRun("nested behavior")

		Run(mockT, func(c *TestContext) {
			something := c.Let(func(c *TestCase) string { return "first" })

			c.It("behavior", func(c *TestCase) {
				testhelp.AssertEqual(t, "first", c.Get(something))
			})

			c.Describe("nested", func(c *TestContext) {
				something := c.Let(func(c *TestCase) string { return "second" })

				c.It("behavior", func(c *TestCase) {
					testhelp.AssertEqual(t, "second", c.Get(something))
				})
			})
		})
	})

	t.Run("caching", func(t *testing.T) {
		mockT := testhelp.NewTestingTMock(t)
		mockT.ExpectRun("behavior")

		calls := 0
		Run(mockT, func(c *TestContext) {
			something := c.Let(func(c *TestCase) string {
				calls++
				return "first"
			})

			c.It("behavior", func(c *TestCase) {
				testhelp.AssertEqual(t, "first", c.Get(something))
				testhelp.AssertEqual(t, "first", c.Get(something))
			})
		})

		testhelp.AssertEqual(t, 1, calls)
	})
}

func TestSet(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		mockT := testhelp.NewTestingTMock(t)
		mockT.ExpectRun("behavior")
		mockT.ExpectRun("when second behavior")

		Run(mockT, func(c *TestContext) {
			something := c.Let(func(c *TestCase) string { return "first" })
			somethingElse := c.Let(func(c *TestCase) string { return c.Get(something) + " thing" })

			c.It("behavior", func(c *TestCase) {
				testhelp.AssertEqual(t, "first thing", c.Get(somethingElse))
			})

			c.Context("when second", func(c *TestContext) {
				c.Set(something, func(c *TestCase) string { return "second" })

				c.It("behavior", func(c *TestCase) {
					testhelp.AssertEqual(t, "second thing", c.Get(somethingElse))
				})
			})
		})
	})

	t.Run("caching", func(t *testing.T) {
		mockT := testhelp.NewTestingTMock(t)
		mockT.ExpectRun("behavior")
		mockT.ExpectRun("when second behavior")

		firstCalls := 0
		thingCalls := 0
		secondCalls := 0

		Run(mockT, func(c *TestContext) {
			something := c.Let(func(c *TestCase) string {
				firstCalls++
				return "first"
			})

			somethingElse := c.Let(func(c *TestCase) string {
				thingCalls++
				return c.Get(something) + " thing"
			})

			c.It("behavior", func(c *TestCase) {
				testhelp.AssertEqual(t, "first thing", c.Get(somethingElse))
				testhelp.AssertEqual(t, "first thing", c.Get(somethingElse)) // Should be cached
			})

			c.Context("when second", func(c *TestContext) {
				c.Set(something, func(c *TestCase) string {
					secondCalls++
					return "second"
				})

				c.It("behavior", func(c *TestCase) {
					testhelp.AssertEqual(t, "second thing", c.Get(somethingElse))
					testhelp.AssertEqual(t, "second thing", c.Get(somethingElse)) // Should be cached
				})
			})
		})

		testhelp.AssertEqual(t, 1, firstCalls)
		testhelp.AssertEqual(t, 2, thingCalls)
		testhelp.AssertEqual(t, 1, secondCalls)
	})
}

func TestContext_findLet(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		context := &TestContext{
			lets: map[string]letFunc{
				"some-let": func(c *TestCase) any { return "child value" },
			},
			parent: &TestContext{
				lets: map[string]letFunc{
					"some-let": func(c *TestCase) any { return "parent value" },
				},
			},
		}

		testhelp.AssertEqual(t, "child value", context.findLet("some-let")(nil))
	})

	t.Run("found in parent", func(t *testing.T) {
		context := &TestContext{
			lets: make(map[string]letFunc),
			parent: &TestContext{
				lets: map[string]letFunc{
					"some-let": func(c *TestCase) any { return "parent value" },
				},
			},
		}

		testhelp.AssertEqual(t, "parent value", context.findLet("some-let")(nil))
	})

	t.Run("undefined", func(t *testing.T) {
		context := &TestContext{
			lets:   make(map[string]letFunc),
			parent: &TestContext{lets: make(map[string]letFunc)},
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
