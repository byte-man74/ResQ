package service

import (
	"fmt"
	"example.com/resq/server/internal/models"
	"example.com/resq/server/internal/repository"
)



type UserService interface {
	CreateUser(user *models.User, registrationLocationInformation *models.UserRegistrationLocation, location *models.Location) (*models.User, error);
	UpdateUserSocialInformation(socialInformation *models.UserSocialInformation) (*models.UserSocialInformation, error)
}


type userService struct {
    repo repository.UserRepository
}



func NewUserService (repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}


func (u *userService) CreateUser (user *models.User, registrationLocation *models.UserRegistrationLocation, location *models.Location) (*models.User, error) {
    createdUser, err := u.repo.CreateUser(
        user, registrationLocation, location,
    )

    if err != nil {
        return nil, fmt.Errorf("unable to create new user account: %w", err)
    }

    return createdUser, nil
}


func (u *userService) UpdateUserSocialInformation(socialInformation *models.UserSocialInformation) (*models.UserSocialInformation, error) {
    updatedInformation, err := u.repo.UpdateUserSocialInformation(socialInformation)

    if err != nil {
        return nil, fmt.Errorf("unable to update user information %w", err)
    }


    return updatedInformation, nil
}
