package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"resq.com/resq/server/internal/routes"
)
var Router *gin.Engine

func InitRouter() {
	Router = gin.New()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())



	routes.UserRoutes(Router, DB)
	log.Println("Router initialized")

}
