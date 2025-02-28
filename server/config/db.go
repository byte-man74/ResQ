package config

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"resq.com/resq/server/models"
)

var DB *gorm.DB

func InitDB() {
	dsn := GetEnv("DATABASE_URL", "xx")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("\033[31mfailed to connect to database: %v\033[0m", err)
	}

	DB = db
	log.Printf("\033[32mDatabase connected successfully\033[0m")
	RunMigrations()
}


func RunMigrations() {
	err := DB.AutoMigrate(models.Models...)
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	log.Println("Database migrations applied successfully")
}
