package user

import (
	"resq/internal/infra/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.Engine, db *gorm.DB) {
	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	userController := NewUserController(userService)

	users := router.Group("users")

	{
		users.POST("/create", userController.CreateUser)
		users.POST("/login", userController.AuthorizeUser)

		users.Use(middleware.AuthMiddleware())
		{
			users.GET("/profile", userController.GetUserProfileInformation)
		}
	}
}
