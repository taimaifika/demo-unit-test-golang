package simplemocking_test

import (
	"demo-unit-test/10-mocking/simplemocking"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) MockTopla(numbers []int) (int, error) {
	// Called informs the mock object that a method has been called and takes a series of values to return
	args := mock.Called(numbers)
	result := args.Get(0)

	return result.(int), args.Error(1)
	// Or
	// return args.Int(0), args.Error(1)
	// Other options: args.Int(0) args.Bool(1) args.String(2)
}

func TestMockTopla(t *testing.T) {

	// Let's create a copy of our test object
	mockRepo := new(MockRepository)

	// setup expectations
	mockRepo.On("MockTopla", []int{2, 3}).Return(5, nil)

	// Call the code we are testing
	testMatematik := simplemocking.Matematik(mockRepo)

	sonuc, err := testMatematik.MockTopla([]int{2, 3})

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)

	assert.Equal(t, 5, sonuc)
	assert.Nil(t, err)
}
