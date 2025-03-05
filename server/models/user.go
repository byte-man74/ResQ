package models




type User struct {
	Base
	Email     string    `gorm:"not null unique;index" json:"email" binding:"required"`
	FirstName string    `gorm:"not null" json:"first_name" binding:"required"`
	LastName  string    `gorm:"not null" json:"last_name" binding:"required"`
	Nin       string    `gorm:"not null" json:"nin" binding:"required"`
	Password  string    `gorm:"not null" json:"password" binding:"required"`
}

type UserRegistrationLocation struct {
	Base
	Location   Location `gorm:"foreignKey:LocationID;constraint:OnDelete:CASCADE" json:"location" binding:"required"`
	LocationID string   `gorm:"not null;index" json:"location_id" binding:"required"`
	DeviceModel string    `gorm:"not null" json:"device_model" binding:"required"`
	UserID      string    `gorm:"not null;index" json:"user_id"`
	UserIP 		string   `gorm:"not null" json:"user_ip"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
}


type UserSocialInformation struct {
	Base
	UserID    string                `gorm:"not null;index" json:"user_id"`
	User      User                  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	Images    []UserHeadShotImages  `gorm:"foreignKey:UserSocialInformationID" json:"images"`
}

type UserHeadShotImages struct {
	Base
	UserSocialInformationID string    `gorm:"not null;index" json:"user_social_information_id"`
	UserID                  string    `gorm:"not null;index" json:"user_id"`
	User                    User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	URL                     string    `gorm:"not null" json:"url"`
}
