package biz_test

import (
	"demo-unit-test/10-mocking/services/groupuser_user/biz"
	"demo-unit-test/10-mocking/services/groupuser_user/entity"
	"demo-unit-test/10-mocking/services/groupuser_user/repository"
	"testing"
)

type userRepoMock struct {
	repository.UserRepo
	FakeCreate func(user *entity.User) error
}

func (m *userRepoMock) Create(user *entity.User) error {
	return m.FakeCreate(user)
}

type groupUserRepoMock struct {
	repository.GroupUserRepo
	FakeCreate  func(group *entity.GroupUser) error
	FakeAddUser func(groupId, userId int) error
}

func (m *groupUserRepoMock) Create(group *entity.GroupUser) error {
	return m.FakeCreate(group)
}

func (m *groupUserRepoMock) AddUser(groupId, userId int) error {
	return m.FakeAddUser(groupId, userId)
}

// Unit tests
func TestUser_Create_1(t *testing.T) {
	userRepo := &userRepoMock{
		FakeCreate: func(user *entity.User) error {
			// mock create user
			user.ID = 1

			return nil
		},
	}

	groupUser := &groupUserRepoMock{
		FakeCreate: func(group *entity.GroupUser) error {
			// mock create group
			group.ID = 1

			return nil
		},
		FakeAddUser: func(groupId, userId int) error {
			// mock add user to group

			return nil
		},
	}

	userBiz := biz.NewUserBiz(userRepo, groupUser)

	userInput := &entity.User{
		Name:    "user1",
		Address: "address1",
	}

	err := userBiz.CreateUser(userInput)
	if err != nil {
		t.Errorf("Expect error is nil, but got %v", err)
	}

	if userInput.ID != 1 {
		t.Errorf("Expect user ID is 1, but got %d", userInput.ID)
	}

	if userInput.Name != "user1" {
		t.Errorf("Expect user name is user1, but got %s", userInput.Name)
	}
}
