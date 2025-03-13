package models

import (
	"resq/pkg/dto"

	"gorm.io/gorm"
)





type User struct {
	gorm.Model
	Email     string `gorm:"not null;unique;index" json:"email" binding:"required"`
	FirstName string `gorm:"not null" json:"first_name" binding:"required"`
	LastName  string `gorm:"not null" json:"last_name" binding:"required"`
	Password  string `gorm:"not null" json:"password" binding:"required"`
}

func (u *User) ToDTO() *dto.UserDTO{
	return &dto.UserDTO{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}
