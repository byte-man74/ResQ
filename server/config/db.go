package config

import (
	"log"
	"resq/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func InitDB () {
	dsn := GetEnv("DB_URL", "")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("\033[31mfailed to connect to database: %v\033[0m", err)
		return
	}
	DB = db
	log.Printf("\033[32mDatabase connected successfully\033[0m")
	RunMigrations()
}

func RunMigrations () {
	err := DB.AutoMigrate(models.Models...)

	if err != nil {
		log.Printf("error creating migrations")
		return
	}

	log.Println("Database migrations applied successfully")
}
