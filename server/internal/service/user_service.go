package service

import (
	"errors"
	"fmt"
	"strings"

	"resq.com/resq/server/dto"
	"resq.com/resq/server/internal/models"
	"resq.com/resq/server/internal/repository"
	"resq.com/resq/server/internal/utils"
)



type UserService interface {
	CreateUser(user *models.User, registrationLocationInformation *models.UserRegistrationLocation, location *models.Location) (*models.User, error)
	UpdateUserSocialInformation(socialInformation *models.UserSocialInformation) (*models.UserSocialInformation, error)
	UpdateUserBasicInformation(user *models.User) (*models.User, error)
	GetFullUserInformation(id string) (*dto.FullUserInformation, error)
	AuthenticateUser(*dto.ILogin) (*dto.IAuthPayload, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) CreateUser(user *models.User, registrationLocation *models.UserRegistrationLocation, location *models.Location) (*models.User, error) {
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
	}

	return updatedUserBasicInformation, nil
}

func (u *userService) GetFullUserInformation(id string) (*dto.FullUserInformation, error) {
	fullInformation, err := u.repo.GetFullUserInformation(id)

	if err != nil {
		return nil, fmt.Errorf("unable to fetch user full information %w", err)
	}

	return fullInformation, nil
}

func (u *userService) AuthenticateUser(login *dto.ILogin) (*dto.IAuthPayload, error) {

	if login.Email == "" {
		return nil, errors.New("user field can't be empty")
	}

	if login.Password == "" {
		return nil, errors.New("password can't be empty")
	}

	login.Email = strings.ToLower(login.Email)
	user, err := u.repo.FindUserByEmail(login.Email)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPassword(user.Password, login.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.Email);
	if err != nil {
		return nil, fmt.Errorf("unable to generate token, %w", err)
	}

	payload := dto.IAuthPayload{
		User:  user,
		Token: token,
	}
	return &payload, nil
}
