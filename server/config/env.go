package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)


func LoadEnv () {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("unable to load env:", err)
	}
	fmt.Println("environment variable connected successfully")
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
