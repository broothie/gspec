package match

import (
	"errors"
	"fmt"

	"github.com/broothie/gspec"
)

// BeError matches errors equivalent to expected according to errors.Is.
func BeError(expected error) gspec.MatcherFunc[error] {
	return func(actual error) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              errors.Is(actual, expected),
			FailureReason:        fmt.Sprintf("expected %v to match error %v", actual, expected),
			NegatedFailureReason: fmt.Sprintf("expected %v not to match error %v", actual, expected),
		}
	}
}

// HaveOccurred matches non-nil errors.
func HaveOccurred() gspec.MatcherFunc[error] {
	return func(actual error) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              actual != nil,
			FailureReason:        "expected an error to have occurred",
			NegatedFailureReason: fmt.Sprintf("expected an error not to have occurred but got %v", actual),
		}
	}
}

// BeErrorType matches errors assignable to T according to errors.As.
func BeErrorType[T error]() gspec.MatcherFunc[error] {
	return func(actual error) gspec.MatchResult {
		var target T
		return gspec.MatchResult{
			IsMatch:              errors.As(actual, &target),
			FailureReason:        fmt.Sprintf("expected %v to match error type %T", actual, target),
			NegatedFailureReason: fmt.Sprintf("expected %v not to match error type %T", actual, target),
		}
	}
}
