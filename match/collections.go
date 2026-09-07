package match

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/broothie/gspec"
)

// BeEmpty matches arrays, channels, maps, slices, and strings with length zero.
func BeEmpty[T any]() gspec.MatcherFunc[T] {
	return func(actual T) gspec.MatchResult {
		length, hasLength := lengthOf(actual)
		return gspec.MatchResult{
			IsMatch:              hasLength && length == 0,
			FailureReason:        fmt.Sprintf("expected %v to be empty", actual),
			NegatedFailureReason: fmt.Sprintf("expected %v not to be empty", actual),
		}
	}
}

// HaveLength matches arrays, channels, maps, slices, and strings with the
// expected length.
func HaveLength[T any](expected int) gspec.MatcherFunc[T] {
	return func(actual T) gspec.MatchResult {
		actualLength, hasLength := lengthOf(actual)
		return gspec.MatchResult{
			IsMatch:              hasLength && actualLength == expected,
			FailureReason:        fmt.Sprintf("expected %v to have length %d but had length %d", actual, expected, actualLength),
			NegatedFailureReason: fmt.Sprintf("expected %v not to have length %d", actual, expected),
		}
	}
}

// Contain matches slices containing all of elements.
func Contain[E comparable](elements ...E) gspec.MatcherFunc[[]E] {
	return func(actual []E) gspec.MatchResult {
		var missingElements []E
		var containedElements []E
		for _, element := range elements {
			if isMissing := slices.Index(actual, element) == -1; isMissing {
				missingElements = append(missingElements, element)
			} else {
				containedElements = append(containedElements, element)
			}
		}

		return gspec.MatchResult{
			IsMatch:              len(missingElements) == 0,
			FailureReason:        fmt.Sprintf("expected %v to include %v", actual, missingElements),
			NegatedFailureReason: fmt.Sprintf("expected %v not to include %v", actual, containedElements),
		}
	}
}

// ConsistOf matches slices containing exactly elements, in any order.
// Duplicate elements are significant.
func ConsistOf[E comparable](elements ...E) gspec.MatcherFunc[[]E] {
	return func(actual []E) gspec.MatchResult {
		remaining := slices.Clone(elements)
		for _, element := range actual {
			index := slices.Index(remaining, element)
			if index == -1 {
				return gspec.MatchResult{
					IsMatch:              false,
					FailureReason:        fmt.Sprintf("expected %v to consist of %v", actual, elements),
					NegatedFailureReason: fmt.Sprintf("expected %v not to consist of %v", actual, elements),
				}
			}

			remaining = slices.Delete(remaining, index, index+1)
		}

		return gspec.MatchResult{
			IsMatch:              len(remaining) == 0,
			FailureReason:        fmt.Sprintf("expected %v to consist of %v", actual, elements),
			NegatedFailureReason: fmt.Sprintf("expected %v not to consist of %v", actual, elements),
		}
	}
}

func lengthOf(value any) (int, bool) {
	reflected := reflect.ValueOf(value)
	if !reflected.IsValid() {
		return 0, false
	}

	switch reflected.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return reflected.Len(), true
	default:
		return 0, false
	}
}
