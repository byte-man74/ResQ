package service

import (
	"fmt"
	"example.com/resq/server/internal/models"
	"example.com/resq/server/internal/repository"
)



type UserService interface {
	CreateUser(user *models.User, registrationLocationInformation *models.UserRegistrationLocation, location *models.Location) (*models.User, error);
	UpdateUserSocialInformation(socialInformation *models.UserSocialInformation) (*models.UserSocialInformation, error);
    UpdateUserBasicInformation(user *models.User) (*models.User, error);
    GetFullUserInformation(id string) (*models.FullUserInformation, error )
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


func (u *userService) UpdateUserBasicInformation(user *models.User) (*models.User, error) {
    updatedUserBasicInformation, err := u.repo.UpdateUser(user)

    if err != nil {
        return nil, fmt.Errorf("unable to update user basic information %w", err)
    };

    return updatedUserBasicInformation, nil
}

func (u *userService) GetFullUserInformation(id string) (*models.FullUserInformation, error ) {
    fullInformation, err := u.repo.GetFullUserInformation(id)

    if err != nil {
        return nil, fmt.Errorf("unable to fetch user full information %w", err)
    }

    return fullInformation, nil
}
