package main

import "resq/config"


func main() {
	config.LoadConfig()

	router := config.Router
	router.HandleMethodNotAllowed = true
	port := config.GetEnv("PORT", ":8000")
	router.Run(port)
}
