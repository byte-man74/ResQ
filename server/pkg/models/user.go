package models





type User struct {
	Base
	Email     string `gorm:"not null;unique;index" json:"email" binding:"required"`
	FirstName string `gorm:"not null" json:"first_name" binding:"required"`
	LastName  string `gorm:"not null" json:"last_name" binding:"required"`
	Password  string `gorm:"not null" json:"password" binding:"required"`
}
