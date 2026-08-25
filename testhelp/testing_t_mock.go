package testhelp

import (
	"slices"
	"strings"
	"testing"
)

type TestingTMock struct {
	t             *testing.T
	expectAllRuns bool
	runs          []run
}

type run struct {
	name   string
	called bool
}

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

func (m *TestingTMock) ExpectAllRuns() {
	m.expectAllRuns = true
}

func (m *TestingTMock) ExpectRun(name string) {
	m.runs = append(m.runs, run{name: name})
}

func (m *TestingTMock) Helper() {}

func (m *TestingTMock) Run(name string, f func(t *testing.T)) bool {
	m.t.Helper()

	if m.expectAllRuns {
		f(new(testing.T))
		return true
	}

	index := slices.IndexFunc(m.runs, func(run run) bool { return run.name == name })
	if index == -1 {
		m.t.Errorf("unexpected call to Run with name %q", name)
		return false
	}

	f(new(testing.T))
	m.runs[index].called = true
	return true
}
