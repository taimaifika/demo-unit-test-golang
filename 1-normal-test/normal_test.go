package normal_test

import (
	normal "demo-unit-test/1-normal-test"
	"testing"
)

func TestAdd(t *testing.T) {
	a, b := 1, 2
	want := 3
	if got := normal.Add(a, b); got != want {
		t.Errorf("Add(%d, %d) = %d, want %d", a, b, got, want)
	}
}
