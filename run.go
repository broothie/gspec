package gspec

import (
	"testing"
)

func Run(t testingT, f ContextFunc) {
	t.Helper()

	context := &Context{lets: make(map[string]letFunc)}
	f(context)

	context.run(t)
}

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

	for _, after := range c.allAfters() {
		after(kase)
	}
}

func (c *Context) runContexts(t testingT) {
	t.Helper()

	for _, context := range c.contexts {
		context.run(t)
	}
}
