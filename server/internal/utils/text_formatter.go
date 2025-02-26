package utils

import "golang.org/x/crypto/bcrypt"


func HashPassword (password string) string {
	// Check for empty password
	if password == "" {
		return ""
	}

	// Generate hash using bcrypt
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hashedBytes)
}
