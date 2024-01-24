package config

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/redis"
)

const (
	redisAddr Key = "REDIS_ADDR"
	redisPass Key = "REDIS_PASS"
	redisDB   Key = "REDIS_DB"
)

func (r *reader) RedisConfig() *redis.Config {
	return &redis.Config{
		Addr:     r.fromEnv(redisAddr),
		Password: r.fromEnv(redisPass),
		DB:       r.fromEnvInt(redisDB),
	}
}
