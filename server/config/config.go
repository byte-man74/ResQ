package config

func LoadConfig() {
	LoadEnv()
	InitDB()
	InitRouter() 
}
