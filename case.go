package gspec

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type CaseFunc func(c *Case)

type Case struct {
	context  *Context
	testingT testingT
	lets     map[string]any
}

func (c *Case) Assert() *assert.Assertions {
	return assert.New(c.testingT)
}

func (c *Case) Require() *require.Assertions {
	return require.New(c.testingT)
}
