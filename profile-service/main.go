package main

import (
	"profile-service/config"
)

func main() {
	config.InitDB()
	config.InitRedis()
	InitRouter()
}
