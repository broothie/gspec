package gspec

// Matcher determines whether a value satisfies an expectation.
type Matcher[A any] interface {
	// Match evaluates actual and returns the result of the comparison.
	Match(actual A) MatchResult
}

// MatcherFunc adapts a function into a Matcher.
type MatcherFunc[A any] func(actual A) MatchResult

// Match applies f to actual.
func (f MatcherFunc[A]) Match(actual A) MatchResult {
	return f(actual)
}

// MatchResult describes whether a matcher accepted a value and why either
// expectation polarity would fail.
type MatchResult struct {
	// IsMatch reports whether the actual value satisfied the matcher.
	IsMatch bool
	// FailureReason explains why a positive expectation failed.
	FailureReason string
	// NegatedFailureReason explains why a negated expectation failed.
	NegatedFailureReason string
}
