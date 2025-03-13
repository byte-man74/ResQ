package user

import (
	"fmt"
	"resq/pkg/dto"
	"resq/pkg/models"

	"gorm.io/gorm"
)




type UserRepository interface {
	CreateUser (*models.User) (*dto.UserDTO, error)
	FindUserByEmail (email string) (*models.User, error)
}


type userRepository struct {
	db *gorm.DB
}

func NewUserRepository (db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}



func (u *userRepository ) CreateUser (user *models.User) (*dto.UserDTO, error) {
	result := u.db.Create(user);
	if result.Error != nil {
		return nil, fmt.Errorf("unable to create user %w", result.Error)
	}
	return user.ToDTO(), nil
}


func (u *userRepository) FindUserByEmail (email string) (*models.User, error) {
	var user models.User
	result := u.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("unable to find user: %w", result.Error)
	}
	return &user, nil
}
