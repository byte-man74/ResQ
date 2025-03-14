package middleware

import (
	"log"
	"net/http"
	"resq/pkg/constants"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func AuthMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{constants.RequestError: "invalid token"})
			ctx.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, jwt.ErrSignatureInvalid
			}
			return constants.JWTSecretKey, nil
		})

		if err != nil {
			log.Printf("JWT Parse error: %v", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{constants.RequestError: "invalid token"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			log.Printf("Invalid token claims or token validation failed")
			ctx.JSON(http.StatusUnauthorized, gin.H{constants.RequestError: "invalid token"})
			ctx.Abort()
			return
		}


		expTime, err := claims.GetExpirationTime()

		if err != nil {
			log.Printf("Error getting expiration time: %v", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{constants.RequestError: "invalid token"})
			ctx.Abort()
			return
		}

		if time.Now().After(expTime.Time) {
			log.Printf("Token expired at: %v", expTime.Time)
			ctx.JSON(http.StatusUnauthorized,  gin.H{constants.RequestError: "invalid token"})
			ctx.Abort()
			return
		}

		// Add claims to context
		ctx.Set("user_id", claims["user_id"])
		ctx.Next()
	}
}
