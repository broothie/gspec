package gspec

import (
	"testing"

	"github.com/broothie/gspec/testhelp"
)

func TestCase_T(t *testing.T) {
	kase := &Case{testingT: t}

	_, ok := any(kase.T()).(*testing.T)
	testhelp.AssertEqual(t, ok, true)
}
