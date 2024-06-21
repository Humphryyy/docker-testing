package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/Humphryyy/docker-testing/api_consumer/global/environment"
	"github.com/redis/go-redis/v9"
)

/* A service interface for Redis */
type RedisService interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd
	Incr(ctx context.Context, key string) *redis.IntCmd
}

/* Redis service singleton */
var rs RedisService

/* Redis service implementation */
type redisService struct {
	client *redis.Client
}

/* Get Redis `GET key` command. It returns redis.Nil error when key does not exist. */
func (r *redisService) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

/* Set Redis `SET key value [expiration]` command. Use expiration for `SETEx`-like behavior. */
func (r *redisService) Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(ctx, key, value, expiration)
}

/* Incr Redis `INCR key` command. If key does not exist creates int of 1, otherwise increments key. */
func (r *redisService) Incr(ctx context.Context, key string) *redis.IntCmd {
	return r.client.Incr(ctx, key)
}

/* Initializes Redis client and service */
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

/* Package wrapper for Redis Get */
func Get(ctx context.Context, key string) (*redis.StringCmd, error) {
	if rs == nil {
		return nil, fmt.Errorf("redis service not initialized")
	}

	return rs.Get(ctx, key), nil
}

/* Package wrapper for Redis Set */
func Set(ctx context.Context, key string, value any, expiration time.Duration) (*redis.StatusCmd, error) {
	if rs == nil {
		return nil, fmt.Errorf("redis service not initialized")
	}

	return rs.Set(ctx, key, value, expiration), nil
}

/* Package wrapper for Redis Incr */
func Incr(ctx context.Context, key string) (*redis.IntCmd, error) {
	if rs == nil {
		return nil, fmt.Errorf("redis service not initialized")
	}

	return rs.Incr(ctx, key), nil
}
