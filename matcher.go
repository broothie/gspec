package gspec

type Matcher[A any] interface {
	Match(actual A) MatchResult
}

type MatcherFunc[A any] func(actual A) MatchResult

func (f MatcherFunc[A]) Match(actual A) MatchResult {
	return f(actual)
}

type MatchResult struct {
	IsMatch              bool
	FailureReason        string
	NegatedFailureReason string
}
