package calculator_test

import (
	"demo-unit-test/11-code-coverage/calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	if got, want := calculator.Add(1, 2), 3; got != want {
		t.Errorf("add method produced wrong result. expected: %d, got: %d", want, got)
	}
}

func TestSubtract(t *testing.T) {
	if got, want := calculator.Subtract(5, 2), 3; got != want {
		t.Errorf("subtract method produced wrong result. expected: %d, got: %d", want, got)
	}
}
