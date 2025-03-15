package models

import (
	"resq/pkg/models"
	"gorm.io/gorm"
)




type Report struct {
	gorm.Model
	Title       string         `gorm:"null" json:"title"` //would be processed by AI agent later during analysis
	Summary     string         `gorm:"null" json:"summary"`
	Category    []ReportCategory `gorm:"foreignKey:CategoryID" json:"categories"`
	IsAnonymous bool          `json:"is_anonymous"`
	CategoryID  uint          `json:"category_id"`
	ReporterID  uint          `json:"reporter_id"`
	Status      string        `gorm:"type:enum('pending','in_progress','resolved','rejected');default:'pending'" json:"status"`
	Reporter    models.User   `gorm:"foreignKey:ReporterID" json:"reporter"`
	Location    ReportLocation `gorm:"foreignKey:LocationID" json:"location"`
	LocationID  uint          `json:"location_id"`
	Files       []ReportFile  `gorm:"foreignKey:ReportID" json:"files"`
	ValidityLevel int         `gorm:"check:validity_level >= 0 AND validity_level <= 5" json:"validity_level"`
}
