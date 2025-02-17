package models



type User struct {
	Base
	Email     string    `gorm:"not null unique" json:"email"`
	FirstName string    `gorm:"not null" json:"first_name"`
	LastName  string    `gorm:"not null" json:"last_name"`
	Nin       string    `gorm:"not null" json:"nin"`
	UserSocialInformation UserSocialInformation `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user_social_information"`
}

type UserRegistrationLocation struct {
	Base
	Location   Location `gorm:"foreignKey:LocationID;constraint:OnDelete:CASCADE" json:"location"`
	LocationID string   `gorm:"not null;index" json:"location_id"`
	DeviceModel string    `gorm:"not null" json:"device_model"`
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

type FullUserInformation struct {
	UserInformation           *User                     `json:"user_information"`
	UserRegistrationLocation  *UserRegistrationLocation `json:"user_registration_location"`
	UserSocialInformation     *UserSocialInformation    `json:"user_social_information"`
}
