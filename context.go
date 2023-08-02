package gspec

import (
	"strings"

	"github.com/samber/lo"
)

type ContextFunc func(c *Context)

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

func (c *Context) Describe(subject string, f ContextFunc) {
	c.Context(subject, f)
}

func (c *Context) Context(context string, f ContextFunc) {
	ctx := &Context{parent: c, name: context, lets: make(map[string]letFunc)}
	c.contexts = append(c.contexts, ctx)

	f(ctx)
}

func (c *Context) It(behavior string, f CaseFunc) {
	c.cases = append(c.cases, caseEntry{
		name: behavior,
		run:  f,
	})
}

func (c *Context) joinNames(strs ...string) string {
	strs = append([]string{c.name}, strs...)

	if c.parent == nil {
		return strings.Join(lo.WithoutEmpty(strs), " ")
	} else {
		return c.parent.joinNames(strs...)
	}
}
