package examples

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/broothie/gspec"
	. "github.com/broothie/gspec/match"
)

type validationError struct {
	field string
}

func (e *validationError) Error() string {
	return fmt.Sprintf("%s is invalid", e.field)
}

func TestExpectations(t *testing.T) {
	gspec.Run(t, func(t *gspec.TestContext) {
		t.Describe("value matchers", func(t *gspec.TestContext) {
			t.It("compares values structurally", func(t *gspec.TestCase) {
				t.Expect([]int{1, 2}).To(Equal([]int{1, 2}))
				t.Expect(map[string]int{"one": 1}).NotTo(Equal(map[string]int{"two": 2}))
			})

			t.It("recognizes nil values", func(t *gspec.TestCase) {
				var pointer *int
				t.Expect(pointer).To(BeNil[*int]())
			})

			t.It("accepts custom predicates", func(t *gspec.TestCase) {
				isEven := func(value int) bool { return value%2 == 0 }
				t.Expect(4).To(Satisfy("be even", isEven))
			})
		})

		t.Describe("ordering matchers", func(t *gspec.TestContext) {
			t.It("compares ordered values", func(t *gspec.TestCase) {
				t.Expect(2).To(BeLessThan(3))
				t.Expect(2).To(BeAtMost(2))
				t.Expect(3).To(BeGreaterThan(2))
				t.Expect(3).To(BeAtLeast(3))
			})

			t.It("compares floating-point values within a tolerance", func(t *gspec.TestCase) {
				t.Expect(3.1415).To(BeCloseTo(3.14, 0.01))
			})
		})

		t.Describe("collection matchers", func(t *gspec.TestContext) {
			t.It("checks collection contents and size", func(t *gspec.TestCase) {
				values := []int{1, 2, 3}

				t.Expect(values).To(Contain(1, 3))
				t.Expect(values).To(ConsistOf(3, 2, 1))
				t.Expect(values).To(HaveLength[[]int](3))
				t.Expect([]int{}).To(BeEmpty[[]int]())
			})
		})

		t.Describe("string matchers", func(t *gspec.TestContext) {
			t.It("checks literal and regular-expression patterns", func(t *gspec.TestCase) {
				value := "hello, world"

				t.Expect(value).To(ContainSubstring("lo, wo"))
				t.Expect(value).To(HavePrefix("hello"))
				t.Expect(value).To(HaveSuffix("world"))
				t.Expect(value).To(MatchRegexp(regexp.MustCompile(`^hello, \w+$`)))
			})
		})

		t.Describe("error matchers", func(t *gspec.TestContext) {
			t.It("checks error presence, identity, and type", func(t *gspec.TestCase) {
				target := errors.New("not found")
				wrapped := fmt.Errorf("lookup failed: %w", target)
				typed := fmt.Errorf("validation failed: %w", &validationError{field: "name"})

				t.Expect(wrapped).To(HaveOccurred())
				t.Expect(wrapped).To(BeError(target))
				t.Expect(typed).To(BeErrorType[*validationError]())
			})
		})

		t.Describe("action matchers", func(t *gspec.TestContext) {
			t.It("observes changes", func(t *gspec.TestCase) {
				value := 1

				t.Expect(func() { value++ }).To(Change(func() int { return value }))
				t.Expect(func() {}).NotTo(Change(func() int { return value }))
			})

			t.It("observes panics", func(t *gspec.TestCase) {
				t.Expect(func() { panic("boom") }).To(Panic())
				t.Expect(func() { panic("boom") }).To(PanicWith("boom"))
				t.Expect(func() {}).NotTo(Panic())
			})
		})
	})
}
