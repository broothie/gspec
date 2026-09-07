package match

import (
	"cmp"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strings"

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

// MatchRegexp matches strings accepted by re.
func MatchRegexp(re *regexp.Regexp) gspec.MatcherFunc[string] {
	return func(actual string) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              re.MatchString(actual),
			FailureReason:        fmt.Sprintf("expected %v to match regular expression %v", actual, re),
			NegatedFailureReason: fmt.Sprintf("expected %v not to match regular expression %v", actual, re),
		}
	}
}

// ContainSubstring matches strings containing substring.
func ContainSubstring(substring string) gspec.MatcherFunc[string] {
	return func(actual string) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              strings.Contains(actual, substring),
			FailureReason:        fmt.Sprintf("expected %q to contain %q", actual, substring),
			NegatedFailureReason: fmt.Sprintf("expected %q not to contain %q", actual, substring),
		}
	}
}

// HavePrefix matches strings beginning with prefix.
func HavePrefix(prefix string) gspec.MatcherFunc[string] {
	return func(actual string) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              strings.HasPrefix(actual, prefix),
			FailureReason:        fmt.Sprintf("expected %q to have prefix %q", actual, prefix),
			NegatedFailureReason: fmt.Sprintf("expected %q not to have prefix %q", actual, prefix),
		}
	}
}

// HaveSuffix matches strings ending with suffix.
func HaveSuffix(suffix string) gspec.MatcherFunc[string] {
	return func(actual string) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              strings.HasSuffix(actual, suffix),
			FailureReason:        fmt.Sprintf("expected %q to have suffix %q", actual, suffix),
			NegatedFailureReason: fmt.Sprintf("expected %q not to have suffix %q", actual, suffix),
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

// BeError matches errors equivalent to expected according to errors.Is.
func BeError(err error) gspec.MatcherFunc[error] {
	return func(actual error) gspec.MatchResult {
		return gspec.MatchResult{
			IsMatch:              errors.Is(actual, err),
			FailureReason:        fmt.Sprintf("expected %v to match error %v", actual, err),
			NegatedFailureReason: fmt.Sprintf("expected %v not to match error %v", actual, err),
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

// Change matches actions that change the value returned by evaluator.
func Change[A any, Evaluator func() A](evaluator Evaluator) gspec.MatcherFunc[func()] {
	return func(action func()) gspec.MatchResult {
		before := evaluator()
		action()
		after := evaluator()

		return gspec.MatchResult{
			IsMatch:              !reflect.DeepEqual(before, after),
			FailureReason:        fmt.Sprintf("expected action to change %v", before),
			NegatedFailureReason: fmt.Sprintf("expected action not to change %v but was changed to %v", before, after),
		}
	}
}

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

func callAndRecover(action func()) (panicked bool, recovered any) {
	completed := false
	func() {
		defer func() { recovered = recover() }()
		action()
		completed = true
	}()

	return !completed, recovered
}
