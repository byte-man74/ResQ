package dto

import "resq.com/resq/server/internal/models"



type FullUserInformation struct {
	UserInformation           *models.User                     `json:"user_information"`
	UserRegistrationLocation  *models.UserRegistrationLocation `json:"user_registration_location"`
	UserSocialInformation     *models.UserSocialInformation    `json:"user_social_information"`
}


type IAuthPayload struct {
	User *models.User;
	Token string;
}

type ILogin struct {
	Email  string
	Password string
}
