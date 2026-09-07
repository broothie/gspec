package gspec

// BeforeEach registers a hook to run before each test case.
func (t *TestContext) BeforeEach(f TestCaseFunc) {
	t.befores = append(t.befores, f)
}

// AfterEach registers a hook to run after each test case.
func (t *TestContext) AfterEach(f TestCaseFunc) {
	t.afters = append(t.afters, f)
}

func (t *TestContext) allBefores() []TestCaseFunc {
	if t.parent == nil {
		return append([]TestCaseFunc(nil), t.befores...)
	}

	befores := t.parent.allBefores()
	return append(befores, t.befores...)
}

func (t *TestContext) allAfters() []TestCaseFunc {
	if t.parent == nil {
		return append([]TestCaseFunc(nil), t.afters...)
	}

	afters := t.parent.allAfters()
	return append(afters, t.afters...)
}
