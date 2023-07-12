package services

import (
	"errors"
	"fmt"
	"math/rand"
	"session-service/config"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

const (
	REDIS_KEY = "session-service-otp-key-%s"
)

func SendOtp(mobileNo string) (bool, error) {
	ctx := context.TODO()
	key := fmt.Sprintf(REDIS_KEY, mobileNo)

	u, err := FetchProfile(mobileNo)

	if u == nil {
		return false, err
	}

	redisClient().Set(ctx, key, ranNumber(), 0)

	log.Info().
		Str("key", key).
		Msg("Send Otp")

	return true, nil
}

func ValidateOtp(mobileNo string, otp int) (bool, error) {
	rc := redisClient()
	ctx := context.TODO()
	key := fmt.Sprintf(REDIS_KEY, mobileNo)

	log.Info().
		Str("Key", key).
		Msg("Validating Otp")

	result, err := rc.Get(ctx, key).Result()

	if err != nil {
		log.Err(err).
			Str("validate_otp_key", key).
			Msg("Error while fetching key")

	}

	n, _ := strconv.Atoi(result)

	if n != otp {
		err = errors.New("Otp validation failed")

		log.Err(err).
			Str("validate_otp_key", key).
			Msg("Otp validation failed")

		return false, err
	}

	rc.Del(ctx, key)

	return true, nil
}

func redisClient() *redis.Client {
	return config.RedisClient()
}

func ranNumber() int {
	return rand.Intn(9999)
}
