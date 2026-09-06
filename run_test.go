package gspec

import (
	"runtime"
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

	t.Run("runs after hooks when the case panics", func(t *testing.T) {
		mockT := testhelp.NewTestingTMock(t)
		afterCalled := false
		c := &Context{afters: []CaseFunc{func(*Case) { afterCalled = true }}}

		func() {
			defer func() { _ = recover() }()
			c.runCase(mockT, caseEntry{run: func(*Case) { panic("boom") }})
		}()

		testhelp.AssertEqual(t, true, afterCalled)
	})

	t.Run("runs after hooks when the case exits", func(t *testing.T) {
		mockT := testhelp.NewTestingTMock(t)
		afterCalled := make(chan struct{}, 1)
		done := make(chan struct{})
		c := &Context{afters: []CaseFunc{func(*Case) { afterCalled <- struct{}{} }}}

		go func() {
			defer close(done)
			c.runCase(mockT, caseEntry{run: func(*Case) { runtime.Goexit() }})
		}()
		<-done

		select {
		case <-afterCalled:
		default:
			t.Error("after hook was not called")
		}
	})
}
