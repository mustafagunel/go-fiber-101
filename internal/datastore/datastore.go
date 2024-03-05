package datastore

import "fiber-tutorial/internal/model"

type IUserDatastore interface {
	CreateUser(u *model.User) (*model.User, error)
	GetUser(u *model.User) (*[]model.User, error)
}
