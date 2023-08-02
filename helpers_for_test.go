package gspec

import (
	"testing"

	"github.com/broothie/gspec/mocks"
	"go.uber.org/mock/gomock"
)

func assignableToTestFunc() gomock.Matcher {
	return gomock.AssignableToTypeOf(func(*testing.T) {})
}

func runFunc(_ string, f func(*testing.T)) bool {
	f(new(testing.T))
	return true
}

func allowTestFuncs(mockT *mocks.MocktestingT, name any) {
	mockT.EXPECT().
		Run(name, assignableToTestFunc()).
		Do(runFunc)
}
