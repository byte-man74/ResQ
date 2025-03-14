package config

import (
	"log"
	"resq/internal/infra/logger"
	"resq/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := GetEnv("DB_URL", "")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.GlobalLogger.Log(logger.ERROR, "Failed to connect to database", map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	DB = db
	logger.GlobalLogger.Log(logger.INFO, "Migrations successful")
	RunMigrations()
}

func RunMigrations() {
	err := DB.AutoMigrate(models.Models...)

	if err != nil {
		logger.GlobalLogger.Log(logger.ERROR, "Failed to connect to database", map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	log.Println("Database migrations applied successfully")
}
