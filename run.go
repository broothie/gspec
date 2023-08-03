// Package gspec provides a collection of test helpers that form a test framework.
package gspec

import (
	"testing"
)

// Run opens a root test group without a label.
func Run(t testingT, f ContextFunc) {
	t.Helper()

	context := &Context{lets: make(map[string]letFunc)}
	f(context)

	context.run(t)
}

// Describe opens a root test group labelled by the provided subject.
func Describe(t testingT, subject string, f ContextFunc) {
	t.Helper()

	Run(t, func(c *Context) { c.Describe(subject, f) })
}

func (c *Context) run(t testingT) {
	t.Helper()

	c.runCases(t)
	c.runContexts(t)
}

func (c *Context) runCases(t testingT) {
	t.Helper()

	for _, entry := range c.cases {
		t.Run(c.joinNames(entry.name), func(t *testing.T) {
			c.runCase(t, entry)
		})
	}
}

func (c *Context) runCase(t testingT, entry caseEntry) {
	t.Helper()

	kase := &Case{context: c, testingT: t, lets: make(map[string]any)}

	for _, before := range c.allBefores() {
		before(kase)
	}

	entry.run(kase)

	for _, after := range reverse(c.allAfters()) {
		after(kase)
	}
}

func (c *Context) runContexts(t testingT) {
	t.Helper()

	for _, context := range c.contexts {
		context.run(t)
	}
}

func reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i := 0; i < (len(slice)+1)/2; i++ {
		result[i] = slice[len(slice)-1-i]
		result[len(slice)-1-i] = slice[i]
	}

	return result
}
