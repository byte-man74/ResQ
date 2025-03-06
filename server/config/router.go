package config

import (
	"log"
	"resq/internal/domain/user"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter () {
	Router = gin.Default()
	user.UserRoutes(Router, DB);
	log.Println("Router initialized")
}
