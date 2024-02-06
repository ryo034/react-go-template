package middleware

import (
	"github.com/rs/cors"
	"net/http"
)

type CORSInfo struct {
	AllowOrigins []string
}

func (c *CORSInfo) isMatch(origin string) bool {
	for _, v := range c.AllowOrigins {
		if v == origin {
			return true
		}
	}
	return false
}

type CorsMiddleware interface {
	Handler(h http.Handler) http.Handler
}

type corsMiddleware struct {
	info    *CORSInfo
	isLocal bool
}

func NewCorsMiddleware(info *CORSInfo, isLocal bool) CorsMiddleware {
	return &corsMiddleware{info, isLocal}
}

func (cm *corsMiddleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := cors.New(cors.Options{
			AllowedOrigins:   cm.info.AllowOrigins,
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Authorization", "Content-Type"},
			AllowCredentials: true,
			Debug:            cm.isLocal,
		})
		c.Handler(h).ServeHTTP(w, r)
	})
}
