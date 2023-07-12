package main

import (
	"log-service/config"
	"log-service/src/events"
)

var eventChan = make(chan events.Event)

func main() {
	config.InitRedis()
	go events.ReceiveMsg(eventChan)
	InitRouter()
}
