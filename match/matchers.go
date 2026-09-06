package match

import (
	"fmt"
	"slices"

	"github.com/broothie/gspec"
)

func Equal[T comparable](expected T) gspec.MatcherFunc[T] {
	return func(actual T) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              actual == expected,
			FailureReason:        fmt.Sprintf("expected %v to equal %v", actual, expected),
			NegatedFailureReason: fmt.Sprintf("expected %v not to equal %v", actual, expected),
		}
	}
}

func Contain[E comparable](element E) gspec.MatcherFunc[[]E] {
	return func(actual []E) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              slices.Index(actual, element) != -1,
			FailureReason:        fmt.Sprintf("expected %v to include %v", actual, element),
			NegatedFailureReason: fmt.Sprintf("expected %v not to include %v", actual, element),
		}
	}
}
