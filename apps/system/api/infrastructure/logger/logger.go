package logger

import (
	"context"
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"time"
)

type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	With(fields ...interface{}) Logger
	LogRequest(ctx context.Context, req *http.Request) time.Time
	LogResponse(ctx context.Context, req *http.Request, st time.Time, sc int)
}

type Config struct {
	TimeFormat string
	UTC        bool
}

func Recovery(zl Logger, stack bool, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			zl.Error("recovered from panic: %v\n%s", err, debug.Stack())
			httpRequest, _ := httputil.DumpRequest(req, false)
			var fields []interface{}
			fields = append(fields, "time", time.Now())
			fields = append(fields, "error", err)
			fields = append(fields, "request", string(httpRequest))
			if stack {
				fields = append(fields, "stack", string(debug.Stack()))
			}
			zl.Error("[Recovery from panic]", fields...)
		}
	}()
}
