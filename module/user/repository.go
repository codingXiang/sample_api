package user

import "sample_api/model"

type Repository interface {
	GetUserList(map[string]interface{}) ([]*model.User, error)
	GetUser(in *model.User) (*model.User, error)
	CreateUser(in *model.User) (*model.User, error)
	UpdateUser(in *model.User) (*model.User, error)
	ModifyUser(in *model.User, column map[string]interface{}) (*model.User, error)
	DeleteUser(in *model.User) error
}
