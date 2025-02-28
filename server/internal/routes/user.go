package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"resq.com/resq/server/internal/controller"
	"resq.com/resq/server/internal/repository"
	"resq.com/resq/server/internal/service"
)

func UserRoutes (router *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	users := router.Group("/users")
	{
		users.POST("/create/", userController.CreateUser)
		users.POST("/login/", userController.AuthenticateUser)
		users.GET("/user-information/:id", userController.GetFullUserInformation)
	}
}
