package middleware

import (
	"bytes"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/logger"
	"io"
	"net/http"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	body   bytes.Buffer
	status int
}

func (lrw *loggingResponseWriter) WriteHeader(status int) {
	lrw.status = status
	lrw.ResponseWriter.WriteHeader(status)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	lrw.body.Write(b)
	return lrw.ResponseWriter.Write(b)
}

func newStatusResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{ResponseWriter: w}
}

type logMiddleware struct {
	zl    logger.Logger
	stack bool
}

func (lm *logMiddleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestBody bytes.Buffer
		requestBodyReader := io.TeeReader(r.Body, &requestBody)
		bodyBytes, _ := io.ReadAll(requestBodyReader)
		r.Body = io.NopCloser(&requestBody)

		lrw := newStatusResponseWriter(w)

		st := lm.zl.LogRequest(r.Context(), r, bodyBytes)
		logger.Recovery(lm.zl, lm.stack, r)
		h.ServeHTTP(lrw, r)
		lm.zl.LogResponse(r.Context(), r, st, lrw.status, lrw.body.String())
	})
}

func NewLogMiddleware(zl logger.Logger, stack bool) Middleware {
	return &logMiddleware{zl, stack}
}
