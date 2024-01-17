package middleware

import (
	"context"
	"github.com/ogen-go/ogen/middleware"
	middl "github.com/ryo034/react-go-template/apps/system/api/infrastructure/middleware"
)

func RequestID() middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		req.SetContext(context.WithValue(req.Context, middl.RequestIDKey, middl.GenRequestID()))
		req.Raw.Header.Set(middl.RequestIDKey, req.Context.Value(middl.RequestIDKey).(string))
		return next(req)
	}
}
