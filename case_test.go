package gspec

import (
	"testing"

	"github.com/broothie/gspec/testhelp"
)

func TestTestCasePromotesTestingTMethods(t *testing.T) {
	parentName := t.Name()

	Run(t, func(t *TestContext) {
		t.It("uses the current subtest", func(t *TestCase) {
			t.Helper()
			testhelp.AssertEqual(t.T, false, t.Name() == parentName)
			testhelp.AssertEqual(t.T, t.Name(), t.T.Name())

			cleanupCalled := false
			t.Cleanup(func() {
				testhelp.AssertEqual(t.T, true, cleanupCalled)
			})
			t.Cleanup(func() { cleanupCalled = true })
		})
	})
}
