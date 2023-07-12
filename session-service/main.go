package main

import (
	"session-service/config"
)

func main() {
	config.InitRedis()
	InitRouter()
}
