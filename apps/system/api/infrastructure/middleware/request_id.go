package middleware

import (
	"context"
	"github.com/rs/xid"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/logger"
	"net/http"
)

func GenRequestID() string {
	return xid.New().String()
}

type requestIDMiddleware struct {
}

func (rim *requestIDMiddleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = GenRequestID()
		}
		ctx := context.WithValue(r.Context(), logger.RequestIDKey, requestID)
		r.Header.Set("X-Request-ID", requestID)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewRequestIDMiddleware() Middleware {
	return &requestIDMiddleware{}
}
