package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"profile-service/config"
	"profile-service/src/models"

	"github.com/rs/zerolog/log"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

const (
	CACHE_USER_KEY = "profile-service-user-%d"
)

func CreateUser(args map[string]interface{}) (*models.User, error) {
	defer handleUserPanic()
	user := models.NewUser(args)
	err := dbClient().Create(user).Error

	if err != nil {
		return nil, err //,# fmt.Errorf("Error %s- %s ", err.Message, err.Detail).(error)
	}

	return user, nil
}

func GetUser(userId int) (*models.User, error) {
	defer handleUserPanic()

	var user models.User

	log.Info().
		Int("user-id", userId).
		Msgf("Service: GetUser - Fetching user details")

	u := fetchUserfromCache(userId)

	if u != nil {
		return u, nil
	}

	err := dbClient().First(&user, userId).Error

	cacheUserDetails(user)

	if err != nil {
		return nil, errors.New("Record Not found") //,# fmt.Errorf("Error %s- %s ", err.Message, err.Detail).(error)
	}

	return &user, nil

}

func dbClient() *gorm.DB {
	return config.DbClient()

}

func redisClient() *redis.Client {
	return config.RedisClient()
}

func handleUserPanic() (*models.User, error) {

	if err := recover(); err != nil {
		err = fmt.Errorf("Db error: %v\n", err)
		log.Print("error", err)
		return nil, err.(error)
	}
	return nil, nil

}

func fetchUserfromCache(userId int) *models.User {

	ctx := context.TODO()

	user := &models.User{}
	key := fmt.Sprintf(CACHE_USER_KEY, userId)

	log.Print("fetchUserfromCache : for key %s", key)

	val, err := redisClient().Get(ctx, key).Result()

	if err != nil {
		log.Print("Error while fetching value from redis", err)
		return nil
	}

	json.Unmarshal([]byte(val), user)

	return user
}

func cacheUserDetails(user models.User) {
	ctx := context.TODO()

	key := fmt.Sprintf(CACHE_USER_KEY, user.Id)

	log.Info().
		Str("cache-key", key).
		Msg("Caching user details")

	data, _ := json.Marshal(user)

	redisClient().Set(ctx, key, data, 0)

}

func GetUserByMobile(mobileNo string) (*models.User, error) {
	defer handleUserPanic()

	var user models.User

	log.Info().
		Str("mobileNo", mobileNo).
		Msgf("Service: GetUser - Fetching user details - by modile")

	err := dbClient().First(&user, "number = ?", mobileNo).Error

	if err != nil {
		return nil, errors.New("Record Not found") //,# fmt.Errorf("Error %s- %s ", err.Message, err.Detail).(error)
	}

	return &user, nil

}
