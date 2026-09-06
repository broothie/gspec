package examples

import (
	"testing"

	"github.com/broothie/gspec"
	"github.com/broothie/gspec/match"
)

func Test_assertions(t *testing.T) {
	gspec.Describe(t, "addition", func(c *gspec.Context) {
		one := c.Let(func(c *gspec.Case) int { return 1 })

		c.It("should sum numbers", func(c *gspec.Case) {
			c.Expect(c.Get(one) + 1).NotTo(match.Equal(1))
			c.Expect(c.Get(one) + 1).To(match.Equal(2))
		})
	})

	gspec.Describe(t, "inclusion", func(c *gspec.Context) {
		ints := c.Let(func(c *gspec.Case) []int { return []int{1, 2, 3} })
		fruits := c.Let(func(c *gspec.Case) []string { return []string{"apple", "banana", "cherry"} })

		c.It("should work", func(c *gspec.Case) {
			c.Expect(c.Get(ints)).To(match.Include(2))
			c.Expect(c.Get(ints)).NotTo(match.Include(4))

			c.Expect(c.Get(fruits)).To(match.Include("banana"))
			c.Expect(c.Get(fruits)).NotTo(match.Include("date"))
		})
	})
}
