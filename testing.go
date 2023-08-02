package gspec

import "testing"

//go:generate mockgen --source testing.go --destination mocks/testing_t.go --package mocks --typed
type testingT interface {
	Helper()
	Run(name string, f func(t *testing.T)) bool
}
