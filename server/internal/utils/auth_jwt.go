package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


const TokenExpiry = time.Hour * 24;
var jwtSecret = []byte("your_secret_key");

func GenerateJWT(userId string,) (string, error){
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp": time.Now().Add(TokenExpiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtSecret, nil
	})

	// If parsing failed, return error
	if err != nil {
		return nil, err
	}

	// Extract claims (payload)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Check expiration (exp claim)
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, errors.New("token has expired")
		}
	}

	// Token is valid, return claims
	return &claims, nil
}
