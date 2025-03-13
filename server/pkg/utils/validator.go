package utils

import (
	"errors"
	"strings"
)


func SanitizeEmail (email string) (string, error) {
	if len(email) < 4 {
		return "", errors.New("email too short")
	}

	if !strings.Contains(email, "@") {
		return "", errors.New("email is invalid")
	}

	return strings.ToLower(email), nil
}


func ValidatePassword (password string ) (string, error){
	if len(password) < 6 {
		return "", errors.New("password is too short")
	}

	return password, nil
}
