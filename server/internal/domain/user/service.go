package user

import (
	"errors"
	"fmt"
	"resq/internal/infra/logger"
	"resq/pkg/constants"
	"resq/pkg/dto"
	"resq/pkg/models"
	"resq/pkg/utils"
)

type UserService interface {
	CreateUser(user *models.User) (*dto.UserDTO, error)
	AuthorizeUser(login *dto.LoginRequestDTO) (string, error)
	GetUserProfileInformation(userId uint) (*dto.UserDTO, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repository: repo}
}

func (u *userService) CreateUser(user *models.User) (*dto.UserDTO, error) {
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
		logger.GlobalLogger.Log(logger.ERROR, fmt.Sprintf("unable to create user: %v", err))
		return nil, errors.New(err.Error())
	}

	return result, nil
}

func (u *userService) AuthorizeUser(login *dto.LoginRequestDTO) (string, error) {
	sanitizedEmail, err := utils.SanitizeEmail(login.Email)
	if err != nil {
		return "", err
	}

	user, err := u.repository.FindUserByEmail(sanitizedEmail)
	if err != nil {
		return "", err
	}

	hasVerifiedPassword := utils.VerifyPassword(login.Password, user.Password)
	if !hasVerifiedPassword {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(fmt.Sprint(user.ID), constants.JWTSecretKey)
	if err != nil {
		return "", errors.New("authorization error")
	}

	return token, nil
}

func (u *userService) GetUserProfileInformation(userId uint) (*dto.UserDTO, error) {
	result, err := u.repository.GetUserProfileInformation(userId)

	if err != nil {
		return nil, err
	}

	return result, nil
}
