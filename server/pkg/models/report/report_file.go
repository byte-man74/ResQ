package models

import "gorm.io/gorm"

type ReportFile struct {
	gorm.Model
	FileType string
	// FileSize
	// MediaUrl

}
