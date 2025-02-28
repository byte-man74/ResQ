package models


type IncidentStatus string

const (
	IncidentStatusPending   IncidentStatus = "pending"
	IncidentStatusActive    IncidentStatus = "active"
	IncidentStatusResolved  IncidentStatus = "resolved"
	IncidentStatusArchived  IncidentStatus = "archived"
)

// Incident represents a reported incident or emergency situation
type Incident struct {
	Base
	UploadedBy    User      `gorm:"foreignKey:UploadedByID;constraint:OnDelete:CASCADE" json:"uploaded_by"`
	UploadedByID  string    `gorm:"not null;index" json:"uploaded_by_id"`
	Title         string    `gorm:"not null;type:varchar(255)" json:"title"`
	Description   string    `gorm:"not null;type:text" json:"description"`
	Status        IncidentStatus `gorm:"not null;default:'pending';type:varchar(50)" json:"status"`
	Categories    []IncidentCategory `gorm:"many2many:incident_categories;" json:"categories"`
	Location      Location `gorm:"foreignKey:LocationID;constraint:OnDelete:CASCADE" json:"location"`
	LocationID    string   `gorm:"not null;index" json:"location_id"`
	Media         []IncidentMedia `gorm:"foreignKey:IncidentID" json:"media"`
}

// IncidentMedia represents media files attached to an incident
type IncidentMedia struct {
	Base
	File        File      `gorm:"foreignKey:FileID;constraint:OnDelete:CASCADE" json:"file"`
	FileID      string    `gorm:"not null;index" json:"file_id"`
	IncidentID  string    `gorm:"not null;index" json:"incident_id"`
	Incident    Incident  `gorm:"foreignKey:IncidentID;constraint:OnDelete:CASCADE" json:"incident"`
}


type IncidentAnalytics struct {
	Base
	IncidentID      string    `gorm:"not null;index" json:"incident_id"`
	Incident        Incident  `gorm:"foreignKey:IncidentID;constraint:OnDelete:CASCADE" json:"incident"`
	ViewCount       int64     `gorm:"not null;default:0" json:"view_count"`
	ResponseTime    int64     `gorm:"not null;default:0" json:"response_time"`
	ResolutionTime  int64     `gorm:"not null;default:0" json:"resolution_time"`
}


type IncidentCategory struct {
	Base
	Name        string `gorm:"not null;type:varchar(100);unique" json:"name"`
	Description string `gorm:"not null;type:text" json:"description"`
	Severity    int    `gorm:"not null;check:severity >= 1 AND severity <= 5" json:"severity"`
	IsActive    bool   `gorm:"not null;default:true" json:"is_active"`
}
