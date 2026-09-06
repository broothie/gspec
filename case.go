package gspec

import (
	"testing"
)

// CaseFunc is the signature of functions passed in (typically anonymously) to *Context.It, *Context.BeforeEach, and
// *Context.AfterEach.
type CaseFunc func(c *Case)

// Case provides access to a test case's *testing.T and lazily evaluated Let values.
type Case struct {
	testingT  testingT
	context   *Context
	letValues map[string]any
}

// T provides the test case's underlying *testing.T.
func (c *Case) T() *testing.T {
	return c.testingT.(*testing.T)
}
