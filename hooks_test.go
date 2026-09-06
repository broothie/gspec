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
