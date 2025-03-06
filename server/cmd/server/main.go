package main

import "resq/config"


func main() {
	config.LoadConfig()

	router := config.Router

	port := config.GetEnv("PORT", ":8000")
	router.Run(port)
}
