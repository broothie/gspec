package gspec

import (
	"fmt"
)

type (
	LetFunc[T any] func(c *Case) T
	letFunc        func(c *Case) any
)

// Let defines a value to be retrieved from within a later-defined test case or hook.
// It returns a function which can be called within a test case or hook to retrieve the value.
//
// Let values are only evaluated if they're called within a test case or hook,
// and the value is cached for the duration of the test case.
//
// Let values can be overwritten in nested groups, but their return type must remain the same.
// When overwriting a Let in this way, the returned function needn't be captured.
// The value will still be registered for the context, even though the function was captured in an outer group.
func Let[T any](c *Context, name string, f LetFunc[T]) LetFunc[T] {
	c.registerLet(name, func(c *Case) any { return f(c) })
	return func(c *Case) T { return c.evaluateLet(name).(T) }
}

func (c *Context) registerLet(name string, f letFunc) {
	c.lets[name] = f
}

func (c *Context) findLet(name string) letFunc {
	if value, ok := c.lets[name]; ok {
		return value
	} else if c.parent != nil {
		return c.parent.findLet(name)
	} else {
		panic(fmt.Sprintf("no let defined with name %q", name))
	}
}

func (c *Case) evaluateLet(name string) any {
	if value, ok := c.lets[name]; ok {
		return value
	}

	let := c.context.findLet(name)
	value := let(c)
	c.lets[name] = value
	return value
}
