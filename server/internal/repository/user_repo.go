package repository

import (
	"errors"
	"fmt"

	"example.com/resq/server/internal/models"
	"example.com/resq/server/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)



type UserRepository interface {
	CreateUser(user *models.User, registrationLocationInformation *models.UserRegistrationLocation, location *models.Location) (*models.User, error);
	UpdateUser(user *models.User)(*models.User, error)
	UpdateUserSocialInformation(socialInformation *models.UserSocialInformation) (*models.UserSocialInformation, error)
	GetFullUserInformation(id string) (*models.FullUserInformation, error )
	FindUserByEmail(email string) (*models.User, error)
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

	if len(user.Password) < 6 {
		return nil, fmt.Errorf("password is too short")
	}
	user.Password = utils.HashPassword(user.Password)


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
		return nil, fmt.Errorf("an error happened when updating user: %w", err)
	}

	return user, nil
}



func (r *userRepository) GetFullUserInformation(id string) (*models.FullUserInformation, error) {
	if id == "" {
		return nil, fmt.Errorf("user ID must be provided to get full information")
	}

	var user models.User
	if err := r.db.
		Preload("UserRegistrationLocation.Location").
		Preload("UserSocialInformation.Images").
		First(&user, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("error getting user information for ID %s: %w", id, err)
	}

	// Manually query associated data using the user ID
	fullInformation := &models.FullUserInformation{
		UserInformation: &user,
	}

	// Efficiently load association
	if err := r.db.Model(&models.UserRegistrationLocation{}).
		Where("user_id = ?", id).
		Preload("Location").
		First(&fullInformation.UserRegistrationLocation).Error; err != nil {
		fullInformation.UserRegistrationLocation = nil // Handle missing data
	}

	if err := r.db.Model(&models.UserSocialInformation{}).
		Where("user_id = ?", id).
		Preload("Images").
		First(&fullInformation.UserSocialInformation).Error; err != nil {
		fullInformation.UserSocialInformation = nil // Handle missing data
	}

	return fullInformation, nil
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


func (r *userRepository) AuthenticateUser (email string, password string) (*models.User, error ) {
	if email == "" {
		return nil, errors.New("user field can't be empty")
	}

	if password == "" {
		return nil, errors.New("password can't be empty")
	}


	user, err := r.FindUserByEmail(email)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}


	return user, nil
}
