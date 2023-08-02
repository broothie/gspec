package gspec

import (
	"strings"
)

// ContextFunc is the signature of functions passed in (typically anonymously) to gspec.Run, gspec.Describe,
// *Context.Describe, and *Context.Context.
type ContextFunc func(c *Context)

// Context provides a handle for test groups to define test cases, nested groups, lets, and hooks.
type Context struct {
	parent *Context
	name   string

	lets     map[string]letFunc
	befores  []CaseFunc
	afters   []CaseFunc
	cases    []caseEntry
	contexts []*Context
}

type caseEntry struct {
	name string
	run  CaseFunc
}

// Describe defines a nested group labelled with the provided subject.
func (c *Context) Describe(subject string, f ContextFunc) {
	c.Context(subject, f)
}

// Context defines a nested group labelled with the provided context.
// Context labels typically begin with "when", "with", or "without".
func (c *Context) Context(context string, f ContextFunc) {
	ctx := &Context{parent: c, name: context, lets: make(map[string]letFunc)}
	c.contexts = append(c.contexts, ctx)

	f(ctx)
}

// It defines a test case labelled with the provided behavior.
func (c *Context) It(behavior string, f CaseFunc) {
	c.cases = append(c.cases, caseEntry{
		name: behavior,
		run:  f,
	})
}

func (c *Context) joinNames(strs ...string) string {
	strs = append([]string{c.name}, strs...)

	if c.parent == nil {
		return strings.TrimSpace(strings.Join(strs, " "))
	} else {
		return c.parent.joinNames(strs...)
	}
}
