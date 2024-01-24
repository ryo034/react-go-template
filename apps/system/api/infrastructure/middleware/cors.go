package middleware

import (
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

func Cors(info *CORSInfo, req *http.Request) {
	origin := req.Header.Get("Origin")
	if info.isMatch(origin) {
		req.Header.Set("Access-Control-Allow-Origin", origin)
	}
	req.Header.Set("Access-Control-Allow-Credentials", "true")
	req.Header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Set-Cookie, Content-Disposition")
	req.Header.Set("Access-Control-Expose-Headers", "Content-Disposition")
	req.Header.Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
}

type CorsMiddleware interface {
	Handler(h http.Handler) http.Handler
}

type corsMiddleware struct {
	info *CORSInfo
}

func (cm *corsMiddleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Cors(cm.info, r)
		h.ServeHTTP(w, r)
	})
}

func NewCorsMiddleware(info *CORSInfo) CorsMiddleware {
	return &corsMiddleware{info}
}
