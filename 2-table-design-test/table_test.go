package tabledesigntest_test

import (
	tabledesigntest "demo-unit-test/2-table-design-test"
	"fmt"
	"testing"
)

var table = []struct {
	a, b, want int
}{
	{1, 2, 3},
	{2, 3, 5},
	{3, 4, 7},
	{4, 5, 9},
}

func TestAdd(t *testing.T) {
	for _, tt := range table {
		if got := tabledesigntest.Add(tt.a, tt.b); got != tt.want {
			t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

// TestAddSubtest
func TestAddSubtest(t *testing.T) {
	for _, tt := range table {
		testName := fmt.Sprintf("%d+%d=%d", tt.a, tt.b, tt.want)
		t.Run(testName, func(t *testing.T) {
			if got := tabledesigntest.Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
