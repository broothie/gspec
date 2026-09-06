package gspec

import (
	"testing"

	"github.com/broothie/gspec/testhelp"
)

func TestContext_BeforeEach(t *testing.T) {
	mockT := testhelp.NewTestingTMock(t)
	mockT.ExpectRun("tests")

	calls := 0
	Run(mockT, func(c *Context) {
		c.BeforeEach(func(c *Case) {
			testhelp.AssertEqual(t, calls, 0)
			calls++
		})

		c.It("tests", func(c *Case) {
			testhelp.AssertEqual(t, calls, 1)
			calls++
		})
	})

	testhelp.AssertEqual(t, calls, 2)
}

func TestContext_AfterEach(t *testing.T) {
	mockT := testhelp.NewTestingTMock(t)
	mockT.ExpectRun("tests")

	calls := 0
	Run(mockT, func(c *Context) {
		c.AfterEach(func(c *Case) {
			testhelp.AssertEqual(t, calls, 1)
			calls++
		})

		c.It("tests", func(c *Case) {
			testhelp.AssertEqual(t, calls, 0)
			calls++
		})
	})

	testhelp.AssertEqual(t, calls, 2)
}

func TestContext_allBeforesDoesNotAliasParentStorage(t *testing.T) {
	called := ""
	parentBefores := make([]CaseFunc, 0, 2)
	parentBefores = append(parentBefores, func(*Case) {})
	parent := &Context{befores: parentBefores}

	first := (&Context{
		parent:  parent,
		befores: []CaseFunc{func(*Case) { called = "first" }},
	}).allBefores()
	_ = (&Context{
		parent:  parent,
		befores: []CaseFunc{func(*Case) { called = "second" }},
	}).allBefores()

	first[1](nil)
	testhelp.AssertEqual(t, "first", called)
}

func TestContext_allAftersDoesNotAliasParentStorage(t *testing.T) {
	called := ""
	parentAfters := make([]CaseFunc, 0, 2)
	parentAfters = append(parentAfters, func(*Case) {})
	parent := &Context{afters: parentAfters}

	first := (&Context{
		parent: parent,
		afters: []CaseFunc{func(*Case) { called = "first" }},
	}).allAfters()
	_ = (&Context{
		parent: parent,
		afters: []CaseFunc{func(*Case) { called = "second" }},
	}).allAfters()

	first[1](nil)
	testhelp.AssertEqual(t, "first", called)
}

func TestContext_AfterEachSupportsParallelNestedCases(t *testing.T) {
	Run(t, func(c *Context) {
		// Keep spare capacity in the root slice, which previously allowed nested
		// contexts to overwrite shared hook storage.
		c.afters = make([]CaseFunc, 0, 4)
		c.AfterEach(func(*Case) {})
		c.AfterEach(func(*Case) {})
		c.AfterEach(func(*Case) {})

		for i := range 20 {
			c.Context(string(rune('a'+i)), func(c *Context) {
				c.AfterEach(func(*Case) {})
				c.It("runs in parallel", func(c *Case) {
					c.T().Parallel()
				})
			})
		}
	})
}
