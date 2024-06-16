package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/Humphryyy/docker-testing/api_consumer/global/environment"
	"github.com/redis/go-redis/v9"
)

type RedisService interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd
	Incr(ctx context.Context, key string) *redis.IntCmd
}

var rs RedisService

type redisService struct {
	client *redis.Client
}

func (r *redisService) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

func (r *redisService) Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(ctx, key, value, expiration)
}

func (r *redisService) Incr(ctx context.Context, key string) *redis.IntCmd {
	return r.client.Incr(ctx, key)
}

func Init() error {
	opts, err := redis.ParseURL(environment.RedisUrl)
	if err != nil {
		return err
	}

	redisClient := redis.NewClient(opts)

	rs = &redisService{
		client: redisClient,
	}

	return nil
}

func Get(ctx context.Context, key string) (*redis.StringCmd, error) {
	if rs == nil {
		return nil, fmt.Errorf("redis service not initialized")
	}

	return rs.Get(ctx, key), nil
}

func Set(ctx context.Context, key string, value any, expiration time.Duration) (*redis.StatusCmd, error) {
	if rs == nil {
		return nil, fmt.Errorf("redis service not initialized")
	}

	return rs.Set(ctx, key, value, expiration), nil
}

func Incr(ctx context.Context, key string) (*redis.IntCmd, error) {
	if rs == nil {
		return nil, fmt.Errorf("redis service not initialized")
	}

	return rs.Incr(ctx, key), nil
}
