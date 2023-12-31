package gspec

// BeforeEach registers a hook to run before each test case.
func (c *Context) BeforeEach(f CaseFunc) {
	c.befores = append(c.befores, f)
}

// AfterEach registers a hook to run after each test case.
func (c *Context) AfterEach(f CaseFunc) {
	c.afters = append(c.afters, f)
}

func (c *Context) allBefores() []CaseFunc {
	if c.parent == nil {
		return c.befores
	} else {
		return append(c.parent.allBefores(), c.befores...)
	}
}

func (c *Context) allAfters() []CaseFunc {
	if c.parent == nil {
		return c.afters
	} else {
		return append(c.parent.allAfters(), c.afters...)
	}
}
