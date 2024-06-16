package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisTestService struct {
	sync.Mutex
	store map[string]any
}

func InitTest() {
	rs = &redisTestService{store: make(map[string]any)}
}

func (r *redisTestService) Get(ctx context.Context, key string) *redis.StringCmd {
	r.Lock()
	defer r.Unlock()

	cmd := redis.NewStringCmd(ctx)

	val, ok := r.store[key]
	if !ok {
		cmd.SetErr(redis.Nil)
	} else {
		cmd.SetVal(fmt.Sprint(val))
	}

	return cmd
}

func (r *redisTestService) Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
	r.Lock()
	defer r.Unlock()

	r.store[key] = value

	cmd := redis.NewStatusCmd(ctx)

	cmd.SetVal("OK")

	return cmd
}

func (r *redisTestService) Incr(ctx context.Context, key string) *redis.IntCmd {
	r.Lock()
	defer r.Unlock()

	cmd := redis.NewIntCmd(ctx)

	val, ok := r.store[key]
	if !ok {
		r.store[key] = int64(1)
		cmd.SetVal(1)
	} else {
		r.store[key] = val.(int64) + 1
		cmd.SetVal(val.(int64) + 1)
	}

	return cmd
}
