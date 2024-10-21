package biz

import "demo-unit-test/10-mocking/services/groupuser_user/repository"

type UserBiz struct {
	userRepo      repository.UserRepo
	groupUserRepo repository.GroupUserRepo
}

func NewUserBiz(userRepo repository.UserRepo, groupUserRepo repository.GroupUserRepo) *UserBiz {
	return &UserBiz{
		userRepo:      userRepo,
		groupUserRepo: groupUserRepo,
	}
}
