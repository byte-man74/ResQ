package models

type Ministry struct {
	Base
	Name    string `gorm:"not null" json:"name"`
	LogoUrl string `gorm:"not null" json:"logo_url"`
}


type Organization struct {
	Base
	Name    string `gorm:"not null" json:"name"`
	LogoUrl string `gorm:"not null" json:"logo_url"`
}
