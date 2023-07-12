package services

import (
	"auth-service/config"
	"auth-service/src/models"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
)

const (
	SessionCacheKey = "auth-service-referesh-key-%s"
	JwtKey          = "kgdhJpDxlMod2uS8HZCaZGKfBYRD957m"
	Alphanum        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func ValidateSession(token string) (*models.User, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtKey), nil
		})

	if err != nil {
		return nil, err
	}

	refreshKey := claims["refKey"].(string)

	u := fetchUserfromCache(refreshKey)

	if u == nil {
		return nil, errors.New("Invalid token")
	}

	return u, nil
}

func fetchUserfromCache(refreshKey string) *models.User {
	ctx := context.TODO()

	user := &models.User{}
	key := fmt.Sprintf(SessionCacheKey, refreshKey)

	log.Print("fetchUserfromCache : for key %s", key)

	val, err := redisClient().Get(ctx, key).Result()

	if err != nil {
		log.Print("Error while fetching value from redis", err)
		return nil
	}

	json.Unmarshal([]byte(val), user)

	return user
}

func CreateSession(mobileNo string, otp int) (interface{}, error) {
	success, err := ValidateOtp(mobileNo, otp)

	if !success {
		return nil, err
	}

	u, err := FetchProfile(mobileNo)

	if err != nil {
		return nil, err
	}

	session := models.Session{
		RefreshKey: newRefreshKey(),
	}

	token := genrateToken(session)
	saveSessiontoCache(u, session.RefreshKey)

	return token, nil

}

func genrateToken(sessionObj models.Session) string {
	key := []byte(JwtKey)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":    "my-auth-server",
			"refKey": sessionObj.RefreshKey,
		})

	token, _ := t.SignedString(key)
	return token
}

func newRefreshKey() string {
	var bytes = make([]byte, 32)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = Alphanum[b%byte(len(Alphanum))]
	}
	return string(bytes)
}

func saveSessiontoCache(u *models.User, refreshKey string) {
	ctx := context.TODO()
	key := fmt.Sprintf(SessionCacheKey, refreshKey)
	data, _ := json.Marshal(u)
	redisClient().Set(ctx, key, data, 0)
}

func redisClient() *redis.Client {
	return config.RedisClient()
}
