package models


type Case struct {
	Base
	UploadedBy    User      `gorm:"foreignKey:UploadedByID;constraint:OnDelete:CASCADE" json:"uploaded_by"`
	UploadedByID  string    `gorm:"not null;index" json:"uploaded_by_id"`
	Title         string    `gorm:"not null" json:"title"`
	Description   string    `gorm:"not null" json:"description"`
	Status        string    `gorm:"not null;default:'pending'" json:"status"`
	Location   Location `gorm:"foreignKey:LocationID;constraint:OnDelete:CASCADE" json:"location"`
	LocationID string   `gorm:"not null;index" json:"location_id"`
	Media         []CaseMedia `gorm:"foreignKey:CaseID" json:"media"`
}


type CaseMedia struct {
	Base
	File     File      `gorm:"foreignKey:FileID;constraint:OnDelete:CASCADE" json:"file"`
	FileID   string    `gorm:"not null;index" json:"file_id"`
	CaseID  string `gorm:"not null;index" json:"case_id"`
	Case    Case   `gorm:"foreignKey:CaseID;constraint:OnDelete:CASCADE" json:"case"`
}


type CaseAnalytics struct {
	
}
