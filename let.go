package gspec

import (
	"fmt"
	"uuid"
)

type (
	// LetFunc computes the value of a Let for a test case.
	LetFunc[T any] func(t *TestCase) T
	letFunc        func(t *TestCase) any
)

// Let is a typed handle to a lazily evaluated value registered on a TestContext.
type Let[T any] struct {
	id string
}

// Let registers a lazily evaluated, per-case value and returns its typed handle.
func (t *TestContext) Let[T any](letFunc LetFunc[T]) Let[T] {
	id := uuid.New().String()
	t.lets[id] = func(t *TestCase) any { return letFunc(t) }

	return Let[T]{id: id}
}

// Set overrides a Let for this context and its descendants.
func (t *TestContext) Set[T any](let Let[T], letFunc LetFunc[T]) {
	t.lets[let.id] = func(t *TestCase) any { return letFunc(t) }
}

func (t *TestContext) findLet(id string) letFunc {
	if value, ok := t.lets[id]; ok {
		return value
	} else if t.parent != nil {
		return t.parent.findLet(id)
	}

	panic(fmt.Sprintf("no Let defined with name %q", id))
}

// Get returns the value of a Let, evaluating and caching it for this test case when first accessed.
func (t *TestCase) Get[T any](let Let[T]) T {
	if value, ok := t.letValues[let.id]; ok {
		return value.(T)
	}

	letFunc := t.context.findLet(let.id)
	value := letFunc(t)
	t.letValues[let.id] = value
	return value.(T)
}
