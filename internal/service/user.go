package service

import (
	"errors"
	"fiber-tutorial/internal/datastore"
	"fiber-tutorial/internal/model"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository datastore.IUserDatastore
}

func NewUserService(mysql_db *datastore.Mysqldb) *UserService {
	return &UserService{
		UserRepository: mysql_db,
	}
}

func (us *UserService) CreateUser(u *model.User) (*model.User, error) {
	var err error
	u.Password, err = hashPassword(u.Password)
	if err != nil {
		return nil, errors.New("an error occured while creating user")
	}
	return us.UserRepository.CreateUser(u)
}
func (us *UserService) GetUser(u *model.User) (*[]model.User, error) {
	return us.UserRepository.GetUser(u)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
