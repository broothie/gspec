package testhelp

import "testing"

func TestTestingTMock_RunUsesRealSubtest(t *testing.T) {
	mock := NewTestingTMock(t)
	mock.ExpectRun("case")

	var subtestName string
	passed := mock.Run("case", func(t *testing.T) {
		subtestName = t.Name()
	})

	AssertEqual(t, true, passed)
	AssertEqual(t, t.Name()+"/case", subtestName)
}
