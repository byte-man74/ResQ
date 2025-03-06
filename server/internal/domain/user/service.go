package user

import (
	"errors"
	"resq/pkg/models"
)


type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService (repo UserRepository) UserService {
	return &userService{repository: repo}
}


func (u *userService) CreateUser (user *models.User) (*models.User, error ) {
	//validate and hash user password here

	result, err := u.repository.CreateUser(user)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return result, nil
}
