package repository

import "demo-unit-test/10-mocking/services/groupuser_user/entity"

type UserRepo interface {
	Read(userId int) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(userId int) error
}

type GroupUserRepo interface {
	Read(groupId int) (*entity.GroupUser, error)
	Create(group *entity.GroupUser) error
	Update(group *entity.GroupUser) error
	Delete(groupId int) error
	ListUser(groupId int) ([]*entity.User, error)
	AddUser(groupId, userId int) error
	RemoveUser(groupId, userId int) error
}
