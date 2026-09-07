package match

import "testing"

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

func TestSatisfy(t *testing.T) {
	even := Satisfy("be even", func(value int) bool { return value%2 == 0 })

	assertMatches(t, even, 2, true)
	assertMatches(t, even, 3, false)
}
