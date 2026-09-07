package match

import (
	"cmp"
	"fmt"

	"github.com/broothie/gspec"
)

// BeLessThan matches ordered values less than max.
func BeLessThan[T cmp.Ordered](max T) gspec.MatcherFunc[T] {
	return func(actual T) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              actual < max,
			FailureReason:        fmt.Sprintf("expected %v to be less than %v", actual, max),
			NegatedFailureReason: fmt.Sprintf("expected %v not to be less than %v", actual, max),
		}
	}
}

// BeAtMost matches ordered values less than or equal to max.
func BeAtMost[T cmp.Ordered](max T) gspec.MatcherFunc[T] {
	return func(actual T) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              actual <= max,
			FailureReason:        fmt.Sprintf("expected %v to be at most %v", actual, max),
			NegatedFailureReason: fmt.Sprintf("expected %v not to be at most %v", actual, max),
		}
	}
}

// BeGreaterThan matches ordered values greater than min.
func BeGreaterThan[T cmp.Ordered](min T) gspec.MatcherFunc[T] {
	return func(actual T) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              actual > min,
			FailureReason:        fmt.Sprintf("expected %v to be greater than %v", actual, min),
			NegatedFailureReason: fmt.Sprintf("expected %v not to be greater than %v", actual, min),
		}
	}
}

// BeAtLeast matches ordered values greater than or equal to min.
func BeAtLeast[T cmp.Ordered](min T) gspec.MatcherFunc[T] {
	return func(actual T) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              actual >= min,
			FailureReason:        fmt.Sprintf("expected %v to be at least %v", actual, min),
			NegatedFailureReason: fmt.Sprintf("expected %v not to be at least %v", actual, min),
		}
	}
}

// BeCloseTo matches floating-point values within tolerance of expected.
func BeCloseTo[T ~float32 | ~float64](expected, tolerance T) gspec.MatcherFunc[T] {
	return func(actual T) gspec.MatchResult {
		difference := actual - expected
		if difference < 0 {
			difference = -difference
		}

		return gspec.MatchResult{
			IsMatch:              actual == expected || difference <= tolerance,
			FailureReason:        fmt.Sprintf("expected %v to be within %v of %v", actual, tolerance, expected),
			NegatedFailureReason: fmt.Sprintf("expected %v not to be within %v of %v", actual, tolerance, expected),
		}
	}
}
