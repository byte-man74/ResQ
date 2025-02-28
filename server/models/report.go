package models

type Report struct {
	Base
	Name        string     `json:"name" gorm:"not null"`
	Description string     `json:"description" gorm:"not null"`
	Status      string     `json:"status" gorm:"not null;default:'pending'"`
	CreatedBy   string     `json:"created_by" gorm:"not null"`
	UpdatedBy   string     `json:"updated_by"`
	Incidents   []Incident `json:"incidents" gorm:"many2many:report_incidents"`
}
