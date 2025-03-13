package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
	"golang.org/x/crypto/argon2"
)

const (
	iterations   = 1         // Number of iterations
	memory  = 64 * 1024 // 64MB memory
	threads = 4         // Number of parallel threads
	keyLen  = 32        // Length of the derived key
	saltLen = 16        // Length of the salt
)

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, saltLen)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func HashPassword(password string) (string, error) {
	validatedPassword, err := ValidatePassword(password)

	if err != nil {
		return "", err
	};

	salt, err := GenerateSalt()
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(validatedPassword), salt, iterations, memory, threads, keyLen)

	hashedPassword := fmt.Sprintf("%s$%s", base64.RawStdEncoding.EncodeToString(salt), base64.RawStdEncoding.EncodeToString(hash))
	return hashedPassword, nil
}

func VerifyPassword(password string, storedHash string) bool {
	parts := strings.Split(storedHash, "$")

	if len(parts) != 2 {
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])

	if err != nil {
		return false
	}

	storedKey, err := base64.RawStdEncoding.DecodeString(parts[1])

	if err != nil {
		return false
	}

	newHash := argon2.IDKey([]byte(password), salt, iterations, memory, threads, keyLen)

	return string(storedKey) == string(newHash)
}
