package examples

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/broothie/gspec"
	"github.com/broothie/gspec/match"
)

func Test_assertions(t *testing.T) {
	gspec.Run(t, func(c *gspec.Context) {
		c.Describe("addition", func(c *gspec.Context) {
			one := c.Let(func(c *gspec.Case) int { return 1 })

			c.It("should sum numbers", func(c *gspec.Case) {
				c.Expect(c.Get(one) + 1).NotTo(match.Equal(1))
				c.Expect(c.Get(one) + 1).To(match.Equal(2))
			})
		})

		c.Describe("regexp", func(c *gspec.Context) {
			subject := c.Let(func(c *gspec.Case) string { return "hello" })
			re := c.Let(func(c *gspec.Case) *regexp.Regexp { return regexp.MustCompile(`ll`) })

			c.It("matches regexps", func(c *gspec.Case) {
				c.Expect(c.Get(subject)).To(match.MatchRegexp(c.Get(re)))
			})
		})

		c.Describe("contain", func(c *gspec.Context) {
			ints := c.Let(func(c *gspec.Case) []int { return []int{1, 2, 3} })
			fruits := c.Let(func(c *gspec.Case) []string { return []string{"apple", "banana", "cherry"} })

			c.It("should work", func(c *gspec.Case) {
				c.Expect(c.Get(ints)).To(match.Contain(2))
				c.Expect(c.Get(ints)).NotTo(match.Contain(4))

				c.Expect(c.Get(fruits)).To(match.Contain("banana"))
				c.Expect(c.Get(fruits)).NotTo(match.Contain("date"))
			})
		})

		c.Describe("errors", func(c *gspec.Context) {
			anError := c.Let(func(c *gspec.Case) error { return errors.New("something happened") })
			theError := c.Let(func(c *gspec.Case) error { return fmt.Errorf("the error: %w", c.Get(anError)) })

			c.It("checks for error matches", func(c *gspec.Case) {
				c.Expect(c.Get(theError)).To(match.BeError(c.Get(anError)))
			})
		})

		c.Describe("change", func(c *gspec.Context) {
			c.It("tests for changes", func(c *gspec.Case) {
				value := 1
				valuePtr := &value

				c.Expect(func() { *valuePtr += 1 }).To(match.Change(func() int { return *valuePtr }))
				c.Expect(func() { *valuePtr += 0 }).NotTo(match.Change(func() int { return *valuePtr }))
			})
		})
	})
}
