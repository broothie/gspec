package examples

import (
	"testing"

	"github.com/broothie/gspec"
	. "github.com/broothie/gspec/match"
)

func somethingThatNeedsTestingT(t *testing.T) {}

func TestTestingT(t *testing.T) {
	gspec.Describe(t, "testing.T", func(t *gspec.TestContext) {
		t.It("embeds the current *testing.T", func(t *gspec.TestCase) {
			t.Helper()
			t.Expect(t.T).NotTo(BeNil[*testing.T]())
			somethingThatNeedsTestingT(t.T)
		})
	})
}
