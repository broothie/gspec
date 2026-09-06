package gspec

import (
	"fmt"
	"testing"

	"github.com/broothie/gspec/testhelp"
)

func TestContext_runCase(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		mockT := testhelp.NewTestingTMock(t)
		called := false

		c := &Context{name: "some context"}
		c.runCase(mockT, caseEntry{
			name: "case",
			run: func(c *Case) {
				called = true
			},
		})

		testhelp.AssertEqual(t, called, true)
	})

	t.Run("with hooks", func(t *testing.T) {
		mockT := testhelp.NewTestingTMock(t)
		calls := 0

		c := &Context{
			name: "some context",
			befores: []CaseFunc{func(c *Case) {
				testhelp.AssertEqual(t, calls, 0)
				calls++
			}},
			afters: []CaseFunc{func(c *Case) {
				testhelp.AssertEqual(t, calls, 2)
				calls++
			}},
		}

		c.runCase(mockT, caseEntry{
			name: "case",
			run: func(c *Case) {
				testhelp.AssertEqual(t, calls, 1)
				calls++
			},
		})

		testhelp.AssertEqual(t, calls, 3)
	})
}

func Test_reverse(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		testhelp.AssertEqual(t, fmt.Sprint(reverse([]int{})), fmt.Sprint([]int{}))
	})

	t.Run("even number of items", func(t *testing.T) {
		testhelp.AssertEqual(t, fmt.Sprint(reverse([]int{1, 2, 3, 4})), fmt.Sprint([]int{4, 3, 2, 1}))
	})

	t.Run("odd number of items", func(t *testing.T) {
		testhelp.AssertEqual(t, fmt.Sprint(reverse([]int{1, 2, 3})), fmt.Sprint([]int{3, 2, 1}))
	})
}
