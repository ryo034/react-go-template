package redis

import (
	r9 "github.com/redis/go-redis/v9"
)

type Config struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisClient(conf *Config) *r9.Client {
	return r9.NewClient(&r9.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})
}
