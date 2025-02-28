package models

import (
	"time"
	"gorm.io/gorm"
)

type Base struct {
	gorm.Model
	ID        string    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}
