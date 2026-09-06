package gspec

import "testing"

type testingT interface {
	Helper()
	Run(name string, f func(t *testing.T)) bool
	Errorf(format string, args ...any)
}
