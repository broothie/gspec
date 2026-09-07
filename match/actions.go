package match

import (
	"fmt"
	"reflect"

	"github.com/broothie/gspec"
)

// Panic matches actions that panic.
func Panic() gspec.MatcherFunc[func()] {
	return func(action func()) gspec.MatchResult {
		panicked, _ := callAndRecover(action)

		return gspec.MatchResult{
			IsMatch:              panicked,
			FailureReason:        "expected action to panic",
			NegatedFailureReason: "expected action not to panic",
		}
	}
}

// PanicWith matches actions that panic with a deeply equal value.
func PanicWith(expected any) gspec.MatcherFunc[func()] {
	return func(action func()) gspec.MatchResult {
		panicked, recovered := callAndRecover(action)
		return gspec.MatchResult{
			IsMatch:              panicked && reflect.DeepEqual(recovered, expected),
			FailureReason:        fmt.Sprintf("expected action to panic with %v but got %v", expected, recovered),
			NegatedFailureReason: fmt.Sprintf("expected action not to panic with %v", expected),
		}
	}
}

func callAndRecover(action func()) (panicked bool, recovered any) {
	completed := false
	func() {
		defer func() { recovered = recover() }()
		action()
		completed = true
	}()

	return !completed, recovered
}
