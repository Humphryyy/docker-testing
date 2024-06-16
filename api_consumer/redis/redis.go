package redis

import (
	"github.com/Humphryyy/docker-testing/api_consumer/global/environment"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func Init() error {
	opts, err := redis.ParseURL(environment.RedisUrl)
	if err != nil {
		return err
	}

	redisClient = redis.NewClient(opts)

	return nil
}

func Client() *redis.Client {
	return redisClient
}
