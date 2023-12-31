package config

import (
	"github.com/rs/cors"
	"strings"
)

const (
	allowOrigin Key = "ALLOW_ORIGIN"
)

func (r *reader) AllowOrigins() []string {
	allowOrigins := strings.Split(r.fromEnv(allowOrigin), ",")
	for i, v := range allowOrigins {
		allowOrigins[i] = strings.TrimSpace(v)
	}
	return allowOrigins
}

func (r *reader) Cors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   r.AllowOrigins(),
		AllowCredentials: !r.IsLocal(),
		Debug:            r.IsLocal(),
		AllowedHeaders:   []string{"*"},
	})
}
