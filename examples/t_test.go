package examples

import (
	"testing"

	"github.com/broothie/gspec"
	. "github.com/broothie/gspec/match"
)

func somethingThatNeedsTestingT(t *testing.T) {}

func TestTestingT(t *testing.T) {
	gspec.Describe(t, ".T", func(c *gspec.Context) {
		c.It("returns a *testing.T", func(c *gspec.Case) {
			c.Expect(c.T()).NotTo(BeNil[*testing.T]())
			somethingThatNeedsTestingT(c.T())
		})
	})
}
