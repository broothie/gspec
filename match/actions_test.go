package match

import "testing"

func TestPanic(t *testing.T) {
	assertMatches(t, Panic(), func() { panic("failed") }, true)
	assertMatches(t, Panic(), func() {}, false)
}

func TestPanicWith(t *testing.T) {
	assertMatches(t, PanicWith([]int{1, 2}), func() { panic([]int{1, 2}) }, true)
	assertMatches(t, PanicWith("expected"), func() { panic("actual") }, false)
	assertMatches(t, PanicWith("expected"), func() {}, false)
}
