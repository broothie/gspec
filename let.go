package gspec

import (
	"fmt"
)

type (
	LetFunc[T any] func(c *Case) T
	letFunc        func(c *Case) any
)

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
