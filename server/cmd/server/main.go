package main

import (
	"resq/config"
	"resq/internal/infra/logger"
)


func main() {
	logger.InitLogger("logs/app.log")
	defer logger.GlobalLogger.Close()
	
	logger.GlobalLogger.Log(logger.INFO, "App started")
	config.LoadConfig()
	router := config.Router
	router.HandleMethodNotAllowed = true
	port := config.GetEnv("PORT", ":8000")
	router.Run(port)
}
