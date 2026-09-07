package gspec

import (
	"testing"
)

// TestCaseFunc is the signature of functions passed in (typically anonymously) to *TestContext.It,
// *TestContext.BeforeEach, and *TestContext.AfterEach.
type TestCaseFunc func(t *TestCase)

// TestCase provides access to a test case's *testing.T and lazily evaluated Let values.
type TestCase struct {
	*testing.T
	context   *TestContext
	letValues map[string]any
}
