package models

import "time"

type File struct {
	Base
	FileUrl     string `gorm:"not null" json:"file_url"`
	FileType    string `gorm:"not null" json:"file_type"`
	FileSummary string `json:"file_summary"`
	FileFormat  string `json:"file_format"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
}


type Location struct {
	Base
	Longitude   float64 `gorm:"not null" json:"longitude"`
	Latitude    float64 `gorm:"not null" json:"latitude"`
	Name        string  `gorm:"not null" json:"name"`
	Description string  `json:"description"`
	Address     string  `gorm:"not null" json:"address"`
	City        string  `gorm:"not null" json:"city"`
	State       string  `gorm:"not null" json:"state"`
	Country     string  `gorm:"not null" json:"country"`
	PostalCode  string  `json:"postal_code"`
}


