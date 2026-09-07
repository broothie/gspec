package match

import (
	"testing"

	"github.com/broothie/gspec"
)

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
