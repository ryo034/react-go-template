package keyvalue

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisDriver struct {
	client *redis.Client
}

func NewRedisDriver(addr string, password string, db int) *RedisDriver {
	return &RedisDriver{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		}),
	}
}

func (r *RedisDriver) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisDriver) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisDriver) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
