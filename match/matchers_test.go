package match

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/broothie/gspec"
)

func TestEqual(t *testing.T) {
	assertMatches(t, Equal([]int{1, 2}), []int{1, 2}, true)
	assertMatches(t, Equal([]int{1, 2}), []int{2, 1}, false)
}

func TestBeNil(t *testing.T) {
	var nilPointer *int
	value := 1

	assertMatches(t, BeNil[*int](), nilPointer, true)
	assertMatches(t, BeNil[*int](), &value, false)
	assertMatches(t, BeNil[int](), 0, false)
}

func TestBeEmpty(t *testing.T) {
	assertMatches(t, BeEmpty[[]int](), []int{}, true)
	assertMatches(t, BeEmpty[string](), "value", false)
	assertMatches(t, BeEmpty[int](), 0, false)
}

func TestHaveLength(t *testing.T) {
	assertMatches(t, HaveLength[[]int](2), []int{1, 2}, true)
	assertMatches(t, HaveLength[string](2), "one", false)
	assertMatches(t, HaveLength[int](0), 0, false)
}

func TestOrderedComparisons(t *testing.T) {
	assertMatches(t, BeLessThan(2), 1, true)
	assertMatches(t, BeLessThan(2), 2, false)
	assertMatches(t, BeAtMost(2), 2, true)
	assertMatches(t, BeAtMost(2), 3, false)
	assertMatches(t, BeGreaterThan(2), 3, true)
	assertMatches(t, BeGreaterThan(2), 2, false)
	assertMatches(t, BeAtLeast(2), 2, true)
	assertMatches(t, BeAtLeast(2), 1, false)
}

func TestSatisfy(t *testing.T) {
	even := Satisfy("be even", func(value int) bool { return value%2 == 0 })

	assertMatches(t, even, 2, true)
	assertMatches(t, even, 3, false)
}

func TestContain(t *testing.T) {
	assertMatches(t, Contain(1, 2), []int{1, 2, 3}, true)
	assertMatches(t, Contain(1, 4), []int{1, 2, 3}, false)
}

func TestConsistOf(t *testing.T) {
	assertMatches(t, ConsistOf(1, 2, 2), []int{2, 1, 2}, true)
	assertMatches(t, ConsistOf(1, 2, 2), []int{1, 2}, false)
	assertMatches(t, ConsistOf(1, 2), []int{1, 2, 3}, false)
}

func TestMatchRegexp(t *testing.T) {
	matcher := MatchRegexp(regexp.MustCompile(`^foo`))

	assertMatches(t, matcher, "foobar", true)
	assertMatches(t, matcher, "barfoo", false)
}

func TestStringMatchers(t *testing.T) {
	assertMatches(t, ContainSubstring("oba"), "foobar", true)
	assertMatches(t, ContainSubstring("baz"), "foobar", false)
	assertMatches(t, HavePrefix("foo"), "foobar", true)
	assertMatches(t, HavePrefix("bar"), "foobar", false)
	assertMatches(t, HaveSuffix("bar"), "foobar", true)
	assertMatches(t, HaveSuffix("foo"), "foobar", false)
}

func TestBeCloseTo(t *testing.T) {
	assertMatches(t, BeCloseTo(1.0, 0.1), 1.05, true)
	assertMatches(t, BeCloseTo(1.0, 0.1), 1.2, false)
}

func TestBeError(t *testing.T) {
	target := errors.New("target")
	wrapped := fmt.Errorf("wrapped: %w", target)

	assertMatches(t, BeError(target), wrapped, true)
	assertMatches(t, BeError(target), errors.New("other"), false)
}

func TestHaveOccurred(t *testing.T) {
	assertMatches(t, HaveOccurred(), errors.New("failed"), true)
	assertMatches(t, HaveOccurred(), nil, false)
}

type exampleError struct{}

func (*exampleError) Error() string { return "example" }

func TestBeErrorType(t *testing.T) {
	wrapped := fmt.Errorf("wrapped: %w", &exampleError{})

	assertMatches(t, BeErrorType[*exampleError](), wrapped, true)
	assertMatches(t, BeErrorType[*exampleError](), errors.New("other"), false)
}

func TestChange(t *testing.T) {
	value := []int{1}
	matcher := Change(func() []int { return value })

	assertMatches(t, matcher, func() { value = append(value, 2) }, true)
	assertMatches(t, matcher, func() {}, false)
}

func TestPanic(t *testing.T) {
	assertMatches(t, Panic(), func() { panic("failed") }, true)
	assertMatches(t, Panic(), func() {}, false)
}

func TestPanicWith(t *testing.T) {
	assertMatches(t, PanicWith([]int{1, 2}), func() { panic([]int{1, 2}) }, true)
	assertMatches(t, PanicWith("expected"), func() { panic("actual") }, false)
	assertMatches(t, PanicWith("expected"), func() {}, false)
}

func assertMatches[A any](t *testing.T, matcher gspec.Matcher[A], actual A, expected bool) {
	t.Helper()

	result := matcher.Match(actual)
	if result.IsMatch != expected {
		t.Errorf("expected match to be %t but got %t", expected, result.IsMatch)
	}
	if result.FailureReason == "" {
		t.Error("expected a failure reason")
	}
	if result.NegatedFailureReason == "" {
		t.Error("expected a negated failure reason")
	}
}
