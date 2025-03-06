package user

import (
	"fmt"
	"resq/pkg/models"

	"gorm.io/gorm"
)




type UserRepository interface {
	CreateUser (*models.User) (*models.User, error)
}


type userRepository struct {
	db *gorm.DB
}

func NewUserRepository (db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}



func (u *userRepository ) CreateUser (user *models.User) (*models.User, error) {
	result := u.db.Create(user);
	if result.Error != nil {
		return nil, fmt.Errorf("unable to create user %w", result.Error)
	}

	return user, nil
}
