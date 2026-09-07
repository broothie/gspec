package match

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/broothie/gspec"
)

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
