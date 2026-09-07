package gspec

import (
	"testing"

	"github.com/broothie/gspec/testhelp"
)

func TestContext_BeforeEach(t *testing.T) {
	mockT := testhelp.NewTestingTMock(t)
	mockT.ExpectRun("tests")

	calls := 0
	Run(mockT, func(c *TestContext) {
		c.BeforeEach(func(c *TestCase) {
			testhelp.AssertEqual(t, calls, 0)
			calls++
		})

		c.It("tests", func(c *TestCase) {
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
	Run(mockT, func(c *TestContext) {
		c.AfterEach(func(c *TestCase) {
			testhelp.AssertEqual(t, calls, 1)
			calls++
		})

		c.It("tests", func(c *TestCase) {
			testhelp.AssertEqual(t, calls, 0)
			calls++
		})
	})

	testhelp.AssertEqual(t, calls, 2)
}

func TestContext_allBeforesDoesNotAliasParentStorage(t *testing.T) {
	called := ""
	parentBefores := make([]TestCaseFunc, 0, 2)
	parentBefores = append(parentBefores, func(*TestCase) {})
	parent := &TestContext{befores: parentBefores}

	first := (&TestContext{
		parent:  parent,
		befores: []TestCaseFunc{func(*TestCase) { called = "first" }},
	}).allBefores()
	_ = (&TestContext{
		parent:  parent,
		befores: []TestCaseFunc{func(*TestCase) { called = "second" }},
	}).allBefores()

	first[1](nil)
	testhelp.AssertEqual(t, "first", called)
}

func TestContext_allAftersDoesNotAliasParentStorage(t *testing.T) {
	called := ""
	parentAfters := make([]TestCaseFunc, 0, 2)
	parentAfters = append(parentAfters, func(*TestCase) {})
	parent := &TestContext{afters: parentAfters}

	first := (&TestContext{
		parent: parent,
		afters: []TestCaseFunc{func(*TestCase) { called = "first" }},
	}).allAfters()
	_ = (&TestContext{
		parent: parent,
		afters: []TestCaseFunc{func(*TestCase) { called = "second" }},
	}).allAfters()

	first[1](nil)
	testhelp.AssertEqual(t, "first", called)
}

func TestContext_AfterEachSupportsParallelNestedCases(t *testing.T) {
	Run(t, func(c *TestContext) {
		// Keep spare capacity in the root slice, which previously allowed nested
		// contexts to overwrite shared hook storage.
		c.afters = make([]TestCaseFunc, 0, 4)
		c.AfterEach(func(*TestCase) {})
		c.AfterEach(func(*TestCase) {})
		c.AfterEach(func(*TestCase) {})

		for i := range 20 {
			c.Context(string(rune('a'+i)), func(c *TestContext) {
				c.AfterEach(func(*TestCase) {})
				c.It("runs in parallel", func(t *TestCase) {
					t.Parallel()
				})
			})
		}
	})
}
