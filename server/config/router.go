package config

import (
	"log"
	"resq/internal/domain/user"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()
	user.UserRoutes(Router, DB)
	Router.RedirectTrailingSlash = true

	log.Println("Router initialized")
}
