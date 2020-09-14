package repository

import (
	"sample_api/model"
	"sample_api/module/user"
)

type UserService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) user.Service {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetUserList(data map[string]interface{}) ([]*model.User, error) {
	return u.repo.GetUserList(data)
}

func (u *UserService) GetUser(in *model.User) (*model.User, error) {
	return u.repo.GetUser(in)
}

func (u *UserService) CreateUser(in *model.User) (*model.User, error) {
	return u.repo.CreateUser(in)
}

func (u *UserService) UpdateUser(in *model.User) (*model.User, error) {
	return u.repo.UpdateUser(in)
}

func (u *UserService) ModifyUser(in *model.User, data map[string]interface{}) (*model.User, error) {
	return u.repo.ModifyUser(in, data)
}

func (u *UserService) DeleteUser(in *model.User) error {
	return u.repo.DeleteUser(in)
}
