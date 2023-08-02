package examples

import (
	"testing"

	"github.com/broothie/gspec"
)

func somethingThatNeedsTestingT(t *testing.T) {}

func Test_t(t *testing.T) {
	gspec.Describe(t, ".T", func(c *gspec.Context) {
		c.It("returns a *testing.T", func(c *gspec.Case) {
			somethingThatNeedsTestingT(c.T()) // <-- here
		})
	})
}
