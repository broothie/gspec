package gspec

import (
	"fmt"
	"uuid"
)

type (
	LetFunc[T any] func(c *Case) T
	letFunc        func(c *Case) any
)

type Let[T any] struct {
	id string
}

func (c *Context) Let[T any](letFunc LetFunc[T]) Let[T] {
	id := uuid.New().String()
	c.lets[id] = func(c *Case) any { return letFunc(c) }

	return Let[T]{id: id}
}

func (c *Context) Set[T any](let Let[T], letFunc LetFunc[T]) {
	c.lets[let.id] = func(c *Case) any { return letFunc(c) }
}

func (c *Context) findLet(id string) letFunc {
	if value, ok := c.lets[id]; ok {
		return value
	} else if c.parent != nil {
		return c.parent.findLet(id)
	}

	panic(fmt.Sprintf("no Let defined with name %q", id))
}

func (c *Case) Get[T any](let Let[T]) T {
	if value, ok := c.letValues[let.id]; ok {
		return value.(T)
	}

	letFunc := c.context.findLet(let.id)
	value := letFunc(c)
	c.letValues[let.id] = value
	return value.(T)
}
