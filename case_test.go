package gspec

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCase_Assert(t *testing.T) {
	kase := &Case{testingT: nil}
	assert.IsType(t, &assert.Assertions{}, kase.Assert())
}

func TestCase_Require(t *testing.T) {
	kase := &Case{testingT: nil}
	assert.IsType(t, &require.Assertions{}, kase.Require())
}
