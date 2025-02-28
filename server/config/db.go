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
		log.Fatal("failed to connect to database: ", err)
	}

	DB = db
	log.Println("Database connected successfully")
	RunMigrations()
}


func RunMigrations() {
	err := DB.AutoMigrate(models.Models...)
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	log.Println("Database migrations applied successfully")
}
