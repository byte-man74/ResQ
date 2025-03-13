package user

import (
	"errors"
	"resq/pkg/dto"
	"resq/pkg/models"
	"resq/pkg/utils"
)


type UserService interface {
	CreateUser(user *models.User) (*dto.UserDTO, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService (repo UserRepository) UserService {
	return &userService{repository: repo}
}


func (u *userService) CreateUser (user *models.User) (*dto.UserDTO, error ) {
	//validate and hash user password here
	sanitizedEmail, err := utils.SanitizeEmail(user.Email)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword
	user.Email = sanitizedEmail

	result, err := u.repository.CreateUser(user)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return result, nil
}


// func (u *userService) AuthorizeUser ()
