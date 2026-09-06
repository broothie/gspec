package testhelp

import "testing"

// AssertEqual reports a test failure when expected and actual are not equal.
func AssertEqual[T comparable](t *testing.T, expected, actual T) {
	t.Helper()

	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
