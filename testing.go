package gspec

import (
	"testing"

	"github.com/stretchr/testify/require"
	_ "go.uber.org/mock/mockgen/model"
)

//go:generate mockgen --source testing.go --destination mocks/testing_t.go --package mocks --typed
type testingT interface {
	require.TestingT
	Helper()
	Run(name string, f func(t *testing.T)) bool
}
