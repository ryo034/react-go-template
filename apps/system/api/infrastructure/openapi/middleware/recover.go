package middleware

import (
	"github.com/ogen-go/ogen/middleware"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/logger"
)

func Recovery(zl logger.Logger, stack bool) middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		logger.Recovery(zl, stack, req.Raw)
		return next(req)
	}
}
