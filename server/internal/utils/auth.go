package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strings"

	"golang.org/x/crypto/argon2"
)

//argon parameter
const (
	Memory      = 64 * 1024
	Iterations  = 3
	Parallelism = 2
	SaltLength  = 16
	KeyLength   = 32
)


func HashPassword (password string) (string, error) {

	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	salt := make([]byte, SaltLength)
	_, err :=rand.Read(salt)

	if err != nil {
		return "", err
	}

	hashedPassword := argon2.IDKey([]byte(password), salt, Iterations, Memory, Parallelism, KeyLength)

	saltBase64 := base64.StdEncoding.EncodeToString(salt)
	hashedPasswordBase64 := base64.StdEncoding.EncodeToString(hashedPassword)

	return saltBase64 + "$" + hashedPasswordBase64, nil
}



func CheckPassword(storedHash, password string) bool {
	parts := strings.Split(storedHash, "$")

	if len(parts) < 2 {
		return false
	}

	salt, err := base64.StdEncoding.DecodeString(parts[0])

	if err != nil {
		return false
	}

	storedHashedPassword, err := base64.StdEncoding.DecodeString(parts[1])

	if err != nil {
		return false
	}

	providedPasswordHash := argon2.IDKey([]byte(password), salt, Iterations, Memory, Parallelism, KeyLength)

	return base64.StdEncoding.EncodeToString(providedPasswordHash) == base64.StdEncoding.EncodeToString(storedHashedPassword)
}
