package services

import (
	"time"

	"github.com/petrokulybaba/go-web-framework/configs"
)

func RedisGet(key string) (string, error) {
	return configs.RedisClient.Get(key).Result()
}

func RedisSet(key string, value interface{}) error {
	return configs.RedisClient.Set(key, value, time.Duration(time.Hour * 24 * 30)).Err()
}

func RedisDelete(key string) error {
	return configs.RedisClient.Del(key).Err()
}
