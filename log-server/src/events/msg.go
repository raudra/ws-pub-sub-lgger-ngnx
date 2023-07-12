package events

import (
	"encoding/json"
	"log-service/config"

	"github.com/go-redis/redis/v8"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

type Event struct {
	Service string                 `json:"service"`
	Payload map[string]interface{} `json:"payload"`
}

const (
	REDIS_CHANNEL = "log-event"
)

func ReceiveMsg(eChan chan Event) {
	log.Info().
		Msg("Channel started")

	ctx := context.TODO()
	pubsub := redisClient().Subscribe(ctx, REDIS_CHANNEL)

	// Close the subscription when we are done.
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(ctx)

		if err != nil {
			log.Err(err).
				Msg("Error while receiving message from channgl")
			continue
		}

		e := Event{}
		_ = json.Unmarshal([]byte(msg.Payload), &e)

		eChan <- e

		log.Info().
			Interface("Msg", e).
			Msgf("Received: Successfully")

	}
}

func redisClient() *redis.Client {
	return config.RedisClient()
}
