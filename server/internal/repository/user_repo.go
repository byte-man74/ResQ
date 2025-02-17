package repository

import (
	"fmt"
	"example.com/resq/server/internal/models"
	"gorm.io/gorm"
)



type UserRepository interface {
	CreateUser(user *models.User, registrationLocationInformation *models.UserRegistrationLocation, location *models.Location) (*models.User, error);
	UpdateUser(user *models.User)(*models.User, error)
	UpdateUserSocialInformation(socialInformation *models.UserSocialInformation, headShotImages *models.UserHeadShotImages) (*models.UserSocialInformation, error)
	GetFullUserInformation(userId string) (*models.FullUserInformation, error )
	// DeleteUser(userId string)(error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User, registrationLocationInformation *models.UserRegistrationLocation, location *models.Location ) (*models.User, error) {

	if user == nil {
		return nil, fmt.Errorf("user value cannot be empty")
	}

	if registrationLocationInformation == nil {
		return nil, fmt.Errorf("registration location information cannot be empty")
	}

	if location == nil {
		return nil, fmt.Errorf("location cannot be nil")
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		//create location first
		if err := tx.Create(location).Error; err != nil {
			return fmt.Errorf("error creating location %w", err)
		}

		//create user
		if err := tx.Create(user).Error; err != nil {
			return fmt.Errorf("error creating user %w", err)
		}


		registrationLocationInformation.LocationID = location.ID
		registrationLocationInformation.UserID = user.ID

		if err := tx.Create(registrationLocationInformation).Error; err != nil {
			return fmt.Errorf("error creating registration user %w", err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}


func (r *userRepository ) UpdateUser (user *models.User) (*models.User, error) {
	if user == nil {
		return nil, fmt.Errorf("user information must be provided")
	}


	if err  := r.db.Save(user).Error; err != nil {
		return nil, fmt.Errorf("an error happened when updating user %w", err)
	}

	return user, nil
}
func (r *userRepository) GetFullUserInformation(userId string) (*models.FullUserInformation, error) {
	if userId == "" {
		return nil, fmt.Errorf("user id must be provided to get full information")
	}

	var user models.User
	if err := r.db.
		Preload("UserRegistrationLocation").
		Preload("UserSocialInformation.Images").
		First(&user, "id = ?", userId).Error; err != nil {
		return nil, fmt.Errorf("error getting user information: %w", err)
	}

	// Once the user and related models are loaded via Preload,
	// there's no need for additional queries
	fullInformation := &models.FullUserInformation{
		UserInformation:          &user,
		UserRegistrationLocation: &user.UserRegistrationLocation,  // Already loaded
		UserSocialInformation:    &user.UserSocialInformation,      // Already loaded
	}

	return fullInformation, nil
}

