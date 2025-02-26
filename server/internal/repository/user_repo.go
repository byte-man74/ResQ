package repository

import (
	"fmt"

	"gorm.io/gorm"
	"resq.com/resq/server/internal/models"
	"resq.com/resq/server/internal/utils"
)

type UserRepository interface {
	CreateUser(user *models.User, registrationLocationInformation *models.UserRegistrationLocation, location *models.Location) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	UpdateUserSocialInformation(socialInformation *models.UserSocialInformation) (*models.UserSocialInformation, error)
	GetFullUserInformation(id string) (*models.FullUserInformation, error)
	FindUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User, registrationLocationInformation *models.UserRegistrationLocation, location *models.Location) (*models.User, error) {

	if user == nil {
		return nil, fmt.Errorf("user value cannot be empty")
	}

	if len(user.Password) < 6 {
		return nil, fmt.Errorf("password is too short")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("unable to secure user password: %w", err)
	}
	user.Password = hashedPassword

	if registrationLocationInformation == nil {
		return nil, fmt.Errorf("registration location information cannot be empty")
	}

	if location == nil {
		return nil, fmt.Errorf("location cannot be nil")
	}

	err = r.db.Transaction(func(tx *gorm.DB) error {
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

func (r *userRepository) UpdateUser(user *models.User) (*models.User, error) {
	if user == nil {
		return nil, fmt.Errorf("user information must be provided")
	}

	if err := r.db.Save(user).Error; err != nil {
		return nil, fmt.Errorf("an error happened when updating user: %w", err)
	}

	return user, nil
}

func (r *userRepository) GetFullUserInformation(id string) (*models.FullUserInformation, error) {
	if id == "" {
		return nil, fmt.Errorf("user ID must be provided to get full information")
	}

	var fullInfo models.FullUserInformation

	err := r.db.Preload("UserRegistrationLocation.Location").
		Preload("UserSocialInformation.Images").
		First(&fullInfo, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("error retrieving full user info: %w", err)
	}

	return &fullInfo, nil
}

func (r *userRepository) UpdateUserSocialInformation(
	socialInformation *models.UserSocialInformation,
) (*models.UserSocialInformation, error) {

	// Save the social information and its associated images together
	if err := r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(socialInformation).Error; err != nil {
		return nil, fmt.Errorf("error saving user social information and images: %w", err)
	}

	return socialInformation, nil
}

func (r *userRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, fmt.Errorf("error finding user with email %s: %w", email, err)
	}

	return &user, nil
}
