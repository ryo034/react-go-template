package middleware

import (
	"github.com/ogen-go/ogen/middleware"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/logger"
)

type LogMiddleware interface {
	Provide(zl logger.Logger) middleware.Middleware
}

func Logger(zl logger.Logger) middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		st := zl.LogRequest(req.Context, req.Raw)
		//respでInternalServerErrorなどが返ってくるため、基本的にはerrはnilになる
		resp, _ := next(req)
		zl.LogResponse(req.Context, req.Raw, st, resp)
		return resp, nil
	}
}
