package testhelp

import (
	"slices"
	"strings"
	"testing"
)

// TestingTMock is a test runner that records and validates expected subtests.
type TestingTMock struct {
	t             *testing.T
	expectAllRuns bool
	runs          []run
}

type run struct {
	name   string
	called bool
}

// NewTestingTMock returns a test runner that reports unexpected and missing runs to t.
func NewTestingTMock(t *testing.T) *TestingTMock {
	t.Helper()

	mock := &TestingTMock{t: t}

	t.Cleanup(func() {
		if mock.expectAllRuns {
			return
		}

		var missedRuns []string
		for _, run := range mock.runs {
			if !run.called {
				missedRuns = append(missedRuns, run.name)
			}
		}

		if len(missedRuns) > 0 {
			t.Errorf("missing runs for:\n- %s", strings.Join(missedRuns, "\n- "))
		}
	})

	return mock
}

// ExpectAllRuns allows every run without requiring it to be registered in advance.
func (m *TestingTMock) ExpectAllRuns() {
	m.expectAllRuns = true
}

// ExpectRun registers the name of a run expected by the mock.
func (m *TestingTMock) ExpectRun(name string) {
	m.runs = append(m.runs, run{name: name})
}

// Helper implements the helper method required by gspec's test runner.
func (m *TestingTMock) Helper() {}

// Errorf reports a formatted error to the underlying test.
func (m *TestingTMock) Errorf(format string, args ...any) {
	m.t.Helper()
	m.t.Errorf(format, args...)
}

// Run executes f as a subtest when name is allowed or was registered as an expected run.
func (m *TestingTMock) Run(name string, f func(t *testing.T)) bool {
	m.t.Helper()

	if m.expectAllRuns {
		return m.t.Run(name, f)
	}

	index := slices.IndexFunc(m.runs, func(run run) bool { return run.name == name })
	if index == -1 {
		m.t.Errorf("unexpected call to Run with name %q", name)
		return false
	}

	m.runs[index].called = true
	return m.t.Run(name, f)
}
