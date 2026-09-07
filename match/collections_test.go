package match

import "testing"

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

func TestContain(t *testing.T) {
	assertMatches(t, Contain(1, 2), []int{1, 2, 3}, true)
	assertMatches(t, Contain(1, 4), []int{1, 2, 3}, false)
}

func TestConsistOf(t *testing.T) {
	assertMatches(t, ConsistOf(1, 2, 2), []int{2, 1, 2}, true)
	assertMatches(t, ConsistOf(1, 2, 2), []int{1, 2}, false)
	assertMatches(t, ConsistOf(1, 2), []int{1, 2, 3}, false)
}
