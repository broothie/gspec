package gspec

type ExpectationTarget[A any] struct {
	t      testingT
	Actual A
}

func Expect[A any](t testingT, actual A) *ExpectationTarget[A] {
	t.Helper()

	return &ExpectationTarget[A]{
		t:      t,
		Actual: actual,
	}
}

func (c *Case) Expect[A any](actual A) *ExpectationTarget[A] {
	c.testingT.Helper()

	return Expect(c.T(), actual)
}

func (e *ExpectationTarget[A]) To(match Matcher[A]) {
	e.t.Helper()

	if result := match.Match(e.Actual); !result.IsMatch {
		e.t.Errorf("%s", result.FailureReason)
	}
}

func (e *ExpectationTarget[A]) NotTo(match Matcher[A]) {
	e.t.Helper()

	if result := match.Match(e.Actual); result.IsMatch {
		e.t.Errorf("%s", result.NegatedFailureReason)
	}
}
