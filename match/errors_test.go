package match

import (
	"errors"
	"fmt"
	"testing"
)

func TestBeError(t *testing.T) {
	target := errors.New("target")
	wrapped := fmt.Errorf("wrapped: %w", target)

	assertMatches(t, BeError(target), wrapped, true)
	assertMatches(t, BeError(target), errors.New("other"), false)
}

func TestHaveOccurred(t *testing.T) {
	assertMatches(t, HaveOccurred(), errors.New("failed"), true)
	assertMatches(t, HaveOccurred(), nil, false)
}

type exampleError struct{}

func (*exampleError) Error() string { return "example" }

func TestBeErrorType(t *testing.T) {
	wrapped := fmt.Errorf("wrapped: %w", &exampleError{})

	assertMatches(t, BeErrorType[*exampleError](), wrapped, true)
	assertMatches(t, BeErrorType[*exampleError](), errors.New("other"), false)
}
