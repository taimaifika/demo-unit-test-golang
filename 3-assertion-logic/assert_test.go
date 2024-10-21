package assertionlogic_test

import (
	assertionlogic "demo-unit-test/3-assertion-logic"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAdd Example with testify/assert
func TestAdd(t *testing.T) {
	got, err := assertionlogic.Add(1, 2)

	// testify assert
	assert.Equal(t, 3, got, "Add(1, 2) should return 3")

	assert.NotEqual(t, 4, got, "Add(1, 2) should not return 4")

	assert.Nil(t, err, "Error should be nil")

	err = errors.New("different then nil")
	if assert.NotNil(t, err) {
		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "different then nil", err.Error())
	}
}

// TestAddTable assertion with table test
func TestAddTable(t *testing.T) {
	table := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
		{4, 5, 9},
	}

	for _, tt := range table {
		got, err := assertionlogic.Add(tt.a, tt.b)

		assert.Equal(t, tt.want, got, "Add(%d, %d) should return %d", tt.a, tt.b, tt.want)
		assert.Nil(t, err, "Error should be nil")
	}
}

// TestAddCustomAssertion custom assertion function
func TestAddCustomAssertion(t *testing.T) {
	// create assertion function
	customAssert := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("Test 1+3=4", func(t *testing.T) {
		got, err := assertionlogic.Add(1, 3)
		customAssert(t, got, 4)
		assert.Nil(t, err, "Error should be nil")
	})

	t.Run("Test 2+4=6", func(t *testing.T) {
		got, err := assertionlogic.Add(2, 4)
		customAssert(t, got, 6)
		assert.Nil(t, err, "Error should be nil")
	})
}
