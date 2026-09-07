package match

import (
	"fmt"
	"reflect"

	"github.com/broothie/gspec"
)

// Equal matches values that are deeply equal.
func Equal[T any](expected T) gspec.MatcherFunc[T] {
	return func(actual T) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              reflect.DeepEqual(actual, expected),
			FailureReason:        fmt.Sprintf("expected %v to equal %v", actual, expected),
			NegatedFailureReason: fmt.Sprintf("expected %v not to equal %v", actual, expected),
		}
	}
}

// BeNil matches nil values, including typed nil pointers, maps, slices,
// functions, interfaces, and channels.
func BeNil[T any]() gspec.MatcherFunc[T] {
	return func(actual T) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              isNil(actual),
			FailureReason:        fmt.Sprintf("expected %v to be nil", actual),
			NegatedFailureReason: fmt.Sprintf("expected %v not to be nil", actual),
		}
	}
}

// Satisfy matches values for which predicate returns true. Description should
// complete the phrase "expected value to ..." in a failure message.
func Satisfy[T any](description string, predicate func(T) bool) gspec.MatcherFunc[T] {
	return func(actual T) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              predicate(actual),
			FailureReason:        fmt.Sprintf("expected %v to %s", actual, description),
			NegatedFailureReason: fmt.Sprintf("expected %v not to %s", actual, description),
		}
	}
}

func isNil(value any) bool {
	reflected := reflect.ValueOf(value)
	if !reflected.IsValid() {
		return true
	}

	switch reflected.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return reflected.IsNil()
	default:
		return false
	}
}
