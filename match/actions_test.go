package match

import "testing"

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
