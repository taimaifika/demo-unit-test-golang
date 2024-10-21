package biz

import (
	"demo-unit-test/10-mocking/services/groupuser_user/entity"
	"errors"
)

func (u *UserBiz) CreateUser(user *entity.User) error {
	// validate user
	if user.Name == "" {
		return errors.New("name is required")
	}

	// create user
	err := u.userRepo.Create(user)
	if err != nil {
		return err
	}

	// create group
	defaultGroup := &entity.GroupUser{
		Name:    user.Name, // default, use user name as group name
		Private: true,      // default, private group
	}
	err = u.groupUserRepo.Create(defaultGroup)
	if err != nil {
		return err
	}

	// add user to group
	err = u.groupUserRepo.AddUser(defaultGroup.ID, user.ID)
	if err != nil {
		return err
	}

	return nil
}
