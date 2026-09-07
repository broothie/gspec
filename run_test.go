package gspec

import (
	"runtime"
	"testing"

	"github.com/broothie/gspec/testhelp"
)

func TestContext_runCase(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		called := false

		c := &TestContext{name: "some context"}
		c.runCase(t, caseEntry{
			name: "case",
			run: func(c *TestCase) {
				called = true
			},
		})

		testhelp.AssertEqual(t, called, true)
	})

	t.Run("with hooks", func(t *testing.T) {
		calls := 0

		c := &TestContext{
			name: "some context",
			befores: []TestCaseFunc{func(c *TestCase) {
				testhelp.AssertEqual(t, calls, 0)
				calls++
			}},
			afters: []TestCaseFunc{func(c *TestCase) {
				testhelp.AssertEqual(t, calls, 2)
				calls++
			}},
		}

		c.runCase(t, caseEntry{
			name: "case",
			run: func(c *TestCase) {
				testhelp.AssertEqual(t, calls, 1)
				calls++
			},
		})

		testhelp.AssertEqual(t, calls, 3)
	})

	t.Run("runs after hooks when the case panics", func(t *testing.T) {
		afterCalled := false
		c := &TestContext{afters: []TestCaseFunc{func(*TestCase) { afterCalled = true }}}

		func() {
			defer func() { _ = recover() }()
			c.runCase(t, caseEntry{run: func(*TestCase) { panic("boom") }})
		}()

		testhelp.AssertEqual(t, true, afterCalled)
	})

	t.Run("runs after hooks when the case exits", func(t *testing.T) {
		afterCalled := make(chan struct{}, 1)
		done := make(chan struct{})
		c := &TestContext{afters: []TestCaseFunc{func(*TestCase) { afterCalled <- struct{}{} }}}

		go func() {
			defer close(done)
			c.runCase(t, caseEntry{run: func(*TestCase) { runtime.Goexit() }})
		}()
		<-done

		select {
		case <-afterCalled:
		default:
			t.Error("after hook was not called")
		}
	})
}
