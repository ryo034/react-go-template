package middleware

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/logger"
	"net/http"
)

type statusResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

type LogMiddleware interface {
	Handler(h http.Handler) http.Handler
}

type logMiddleware struct {
	zl    logger.Logger
	stack bool
}

func (hl *logMiddleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		st := hl.zl.LogRequest(r.Context(), r)
		logger.Recovery(hl.zl, hl.stack, r)
		sw := &statusResponseWriter{ResponseWriter: w}
		h.ServeHTTP(sw, r)
		hl.zl.LogResponse(r.Context(), r, st, sw.status)
	})
}

func NewLogMiddleware(zl logger.Logger, stack bool) LogMiddleware {
	return &logMiddleware{zl, stack}
}
