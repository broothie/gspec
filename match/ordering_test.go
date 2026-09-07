package match

import "testing"

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

func TestBeCloseTo(t *testing.T) {
	assertMatches(t, BeCloseTo(1.0, 0.1), 1.05, true)
	assertMatches(t, BeCloseTo(1.0, 0.1), 1.2, false)
}
