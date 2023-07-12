package events

import (
	"context"
	"encoding/json"
	"fmt"
	"session-service/config"

	"github.com/go-redis/redis/v8"
)

const (
	REDIS_CHANNEL = "log-event"
)

type Event struct {
	Service string      `json:"service"`
	Payload interface{} `json:"payload"`
}

func PublishMsg(args interface{}) {
	ctx := context.TODO()

	fmt.Println(args)
	var payload interface{}
	_ = json.Unmarshal([]byte(args.(string)), &payload)

	e := Event{
		Service: "session-service",
		Payload: payload,
	}

	data, _ := json.Marshal(e)

	if err := redisClient().Publish(ctx, REDIS_CHANNEL, data).Err(); err != nil {
		panic(err)
	}

}

func redisClient() *redis.Client {
	return config.RedisClient()
}
