package main

import (
	"auth-service/config"
)

func main() {
	config.InitRedis()
	InitRouter()
}
