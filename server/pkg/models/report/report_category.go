package models

import "gorm.io/gorm"

type ReportCategory struct {
	gorm.Model
	Title string
	Description string 
}
