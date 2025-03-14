package config

import (
	"os"
	"resq/internal/infra/logger"

	"github.com/joho/godotenv"
)


func LoadEnv () {
	err := godotenv.Load()
	if err != nil {
		logger.GlobalLogger.Log(logger.ERROR, "unable to connect to environment variable", map[string]interface{}{
			"error": err.Error(),
		})
	}
	logger.GlobalLogger.Log(logger.INFO, "Environment variable found")
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
