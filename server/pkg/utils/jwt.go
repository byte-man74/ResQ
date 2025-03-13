package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)


type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT (userID string, secret []byte) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}


// ValidateJWT checks if the token is valid and extracts claims
func ValidateJWT(tokenString string, secret string) (*Claims, error) {
	claims := &Claims{}

	// Parse token and validate claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		return nil, err // Invalid token
	}

	return claims, nil
}
