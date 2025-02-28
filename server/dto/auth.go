package dto

import "resq.com/resq/server/models"



type FullUserInformation struct {
	UserInformation           *models.User                     `json:"user_information"`
	UserRegistrationLocation  *models.UserRegistrationLocation `json:"user_registration_location"`
	UserSocialInformation     *models.UserSocialInformation    `json:"user_social_information"`
}


type IAuthPayload struct {
	User *models.User;
	Token string;
}

type ICreateAccount struct 	{
	User                           models.User                      `json:"user"`
	RegistrationLocationInformation models.UserRegistrationLocation `json:"registrationLocation"`
	Location                       models.Location                  `json:"location"`
}

type ILogin struct {
	Email  string
	Password string
}
