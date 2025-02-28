package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"resq.com/resq/server/config"
)

func main() {
	config.LoadConfig()

	router := config.Router

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Method Not Allowed",
		})
	})

	port := config.GetEnv("PORT", ":8000")
	router.Run(port)
}
