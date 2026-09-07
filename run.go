// Package gspec provides a collection of test helpers that form a test framework.
package gspec

import (
	"testing"
)

// Run opens a root test group without a label.
func Run(t testingT, f TestContextFunc) {
	t.Helper()

	context := &TestContext{lets: make(map[string]letFunc)}
	f(context)

	context.run(t)
}

// Describe opens a root test group labelled by the provided subject.
func Describe(t testingT, subject string, f TestContextFunc) {
	t.Helper()

	Run(t, func(t *TestContext) { t.Describe(subject, f) })
}

func (t *TestContext) run(runner testingT) {
	runner.Helper()

	t.runCases(runner)
	t.runContexts(runner)
}

func (t *TestContext) runCases(runner testingT) {
	runner.Helper()

	for _, entry := range t.cases {
		runner.Run(t.joinNames(entry.name), func(testingT *testing.T) {
			t.runCase(testingT, entry)
		})
	}
}

func (t *TestContext) runCase(testingT *testing.T, entry caseEntry) {
	testingT.Helper()

	testCase := &TestCase{
		T:         testingT,
		context:   t,
		letValues: make(map[string]any),
	}

	for _, after := range t.allAfters() {
		defer after(testCase)
	}

	for _, before := range t.allBefores() {
		before(testCase)
	}

	entry.run(testCase)
}

func (t *TestContext) runContexts(runner testingT) {
	runner.Helper()

	for _, context := range t.contexts {
		context.run(runner)
	}
}
