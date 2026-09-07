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
	gspec.Run(t, func(c *gspec.Context) {
		c.Describe("value matchers", func(c *gspec.Context) {
			c.It("compares values structurally", func(c *gspec.Case) {
				c.Expect([]int{1, 2}).To(Equal([]int{1, 2}))
				c.Expect(map[string]int{"one": 1}).NotTo(Equal(map[string]int{"two": 2}))
			})

			c.It("recognizes nil values", func(c *gspec.Case) {
				var pointer *int
				c.Expect(pointer).To(BeNil[*int]())
			})

			c.It("accepts custom predicates", func(c *gspec.Case) {
				isEven := func(value int) bool { return value%2 == 0 }
				c.Expect(4).To(Satisfy("be even", isEven))
			})
		})

		c.Describe("ordering matchers", func(c *gspec.Context) {
			c.It("compares ordered values", func(c *gspec.Case) {
				c.Expect(2).To(BeLessThan(3))
				c.Expect(2).To(BeAtMost(2))
				c.Expect(3).To(BeGreaterThan(2))
				c.Expect(3).To(BeAtLeast(3))
			})

			c.It("compares floating-point values within a tolerance", func(c *gspec.Case) {
				c.Expect(3.1415).To(BeCloseTo(3.14, 0.01))
			})
		})

		c.Describe("collection matchers", func(c *gspec.Context) {
			c.It("checks collection contents and size", func(c *gspec.Case) {
				values := []int{1, 2, 3}

				c.Expect(values).To(Contain(1, 3))
				c.Expect(values).To(ConsistOf(3, 2, 1))
				c.Expect(values).To(HaveLength[[]int](3))
				c.Expect([]int{}).To(BeEmpty[[]int]())
			})
		})

		c.Describe("string matchers", func(c *gspec.Context) {
			c.It("checks literal and regular-expression patterns", func(c *gspec.Case) {
				value := "hello, world"

				c.Expect(value).To(ContainSubstring("lo, wo"))
				c.Expect(value).To(HavePrefix("hello"))
				c.Expect(value).To(HaveSuffix("world"))
				c.Expect(value).To(MatchRegexp(regexp.MustCompile(`^hello, \w+$`)))
			})
		})

		c.Describe("error matchers", func(c *gspec.Context) {
			c.It("checks error presence, identity, and type", func(c *gspec.Case) {
				target := errors.New("not found")
				wrapped := fmt.Errorf("lookup failed: %w", target)
				typed := fmt.Errorf("validation failed: %w", &validationError{field: "name"})

				c.Expect(wrapped).To(HaveOccurred())
				c.Expect(wrapped).To(BeError(target))
				c.Expect(typed).To(BeErrorType[*validationError]())
			})
		})

		c.Describe("action matchers", func(c *gspec.Context) {
			c.It("observes changes", func(c *gspec.Case) {
				value := 1

				c.Expect(func() { value++ }).To(Change(func() int { return value }))
				c.Expect(func() {}).NotTo(Change(func() int { return value }))
			})

			c.It("observes panics", func(c *gspec.Case) {
				c.Expect(func() { panic("boom") }).To(Panic())
				c.Expect(func() { panic("boom") }).To(PanicWith("boom"))
				c.Expect(func() {}).NotTo(Panic())
			})
		})
	})
}
