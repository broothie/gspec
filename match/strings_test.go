package match

import (
	"regexp"
	"testing"
)

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
