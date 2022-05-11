package components

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisTool struct {
	redisClient *redis.Client
	// ctx 可以传递上下文的一些信息，以及控制时间是否到期
	Ctx context.Context
}

func (redisTool *redisTool) connect() {
	redisTool.redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func (redisTool *redisTool) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return redisTool.redisClient.Set(redisTool.Ctx, key, value, expiration)
}

func (redisTool *redisTool) Get(key string) (string, error) {
	return redisTool.redisClient.Get(redisTool.Ctx, key).Result()
}

func NewRedisClient() *redisTool {
	redisTool := &redisTool{}
	redisTool.connect()
	redisTool.Ctx = context.Background()

	return redisTool
}
