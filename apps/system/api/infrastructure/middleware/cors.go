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

func Cors(info *CORSInfo, w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if info.isMatch(origin) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Set-Cookie, Content-Disposition")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
}

type CorsMiddleware interface {
	Handler(h http.Handler) http.Handler
}

type corsMiddleware struct {
	info *CORSInfo
}

func (cm *corsMiddleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Cors(cm.info, w, r)
		//TODO:
		if r.Method == "OPTIONS" {
			// if Preflight Request, return response here
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}
