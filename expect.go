package gspec

// ExpectationTarget holds the actual value to which a matcher is applied.
type ExpectationTarget[A any] struct {
	t testingT
	// Actual is the value being matched.
	Actual A
}

// Expect creates an expectation for actual that reports failures to t.
func Expect[A any](t testingT, actual A) *ExpectationTarget[A] {
	t.Helper()

	return &ExpectationTarget[A]{
		t:      t,
		Actual: actual,
	}
}

// Expect creates an expectation for actual in this test case.
func (c *Case) Expect[A any](actual A) *ExpectationTarget[A] {
	c.testingT.Helper()

	return Expect(c.T(), actual)
}

// To reports a failure unless match accepts the actual value.
func (e *ExpectationTarget[A]) To(match Matcher[A]) {
	e.t.Helper()

	if result := match.Match(e.Actual); !result.IsMatch {
		e.t.Errorf("%s", result.FailureReason)
	}
}

// NotTo reports a failure if match accepts the actual value.
func (e *ExpectationTarget[A]) NotTo(match Matcher[A]) {
	e.t.Helper()

	if result := match.Match(e.Actual); result.IsMatch {
		e.t.Errorf("%s", result.NegatedFailureReason)
	}
}
