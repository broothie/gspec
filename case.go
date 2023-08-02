package gspec

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// CaseFunc is the signature of functions passed in (typically anonymously) to *Context.It, *Context.BeforeEach, and
// *Context.AfterEach.
type CaseFunc func(c *Case)

// Case provides a handle for test cases to make assertions via *Case.Assert.
// It also provides *Case.Require for assertions that immediately fail the test case.
type Case struct {
	context  *Context
	testingT testingT
	lets     map[string]any
}

// T provides the test case's underlying *testing.T.
func (c *Case) T() *testing.T {
	return c.testingT.(*testing.T)
}

// Assert provides a reference to a test case's *assert.Assertions.
func (c *Case) Assert() *assert.Assertions {
	return assert.New(c.T())
}

// Require provides a reference to a test case's *require.Assertions.
func (c *Case) Require() *require.Assertions {
	return require.New(c.T())
}
