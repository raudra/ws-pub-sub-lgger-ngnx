package main

import (
	"otp-service/config"
)

func main() {
	config.InitRedis()
	InitRouter()
}
