package config

import "os"

func LoadConfig() {
	LoadEnv()
	InitDB()
	InitRouter()
}

var JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))
