package gspec

import (
	"strings"
)

// TestContextFunc is the signature of functions passed in (typically anonymously) to gspec.Run, gspec.Describe,
// *TestContext.Describe, and *TestContext.Context.
type TestContextFunc func(t *TestContext)

// TestContext provides a handle for test groups to define test cases, nested groups, lets, and hooks.
type TestContext struct {
	parent *TestContext
	name   string

	lets     map[string]letFunc
	befores  []TestCaseFunc
	afters   []TestCaseFunc
	cases    []caseEntry
	contexts []*TestContext
}

type caseEntry struct {
	name string
	run  TestCaseFunc
}

// Describe defines a nested group labelled with the provided subject.
func (t *TestContext) Describe(subject string, f TestContextFunc) {
	t.Context(subject, f)
}

// Context defines a nested group labelled with the provided context.
// Context labels typically begin with "when", "with", or "without".
func (t *TestContext) Context(context string, f TestContextFunc) {
	ctx := &TestContext{parent: t, name: context, lets: make(map[string]letFunc)}
	t.contexts = append(t.contexts, ctx)

	f(ctx)
}

// It defines a test case labelled with the provided behavior.
func (t *TestContext) It(behavior string, f TestCaseFunc) {
	t.cases = append(t.cases, caseEntry{
		name: behavior,
		run:  f,
	})
}

func (t *TestContext) joinNames(strs ...string) string {
	strs = append([]string{t.name}, strs...)

	if t.parent == nil {
		return strings.TrimSpace(strings.Join(strs, " "))
	}

	return t.parent.joinNames(strs...)
}
