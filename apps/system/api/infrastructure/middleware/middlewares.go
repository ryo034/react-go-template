package middleware

import (
	"net/http"

	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"

	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/injector"

	"github.com/redis/go-redis/v9"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/logger"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
)

type Middlewares interface {
	Global(h http.Handler, conf config.Reader, zl logger.Logger, rds *redis.Client, p core.Provider, di injector.Driver) http.Handler
}

type Middleware interface {
	Handler(h http.Handler) http.Handler
}

type HttpMiddleware func(http.Handler) http.Handler

type mid struct {
	co shared.ContextOperator
}

func applyMiddlewares(h http.Handler, middlewares ...HttpMiddleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func (m *mid) Global(h http.Handler, conf config.Reader, zl logger.Logger, rds *redis.Client, p core.Provider, di injector.Driver) http.Handler {
	return applyMiddlewares(
		h,
		NewRequestIDMiddleware().Handler,
		NewLangMiddleware(conf.DefaultLanguage(), m.co).Handler,
		NewLogMiddleware(zl, conf.IsLocal()).Handler,
		NewCorsMiddleware(&CORSInfo{AllowOrigins: conf.AllowOrigins()}, conf.IsLocal()).Handler,
		NewOtpRateLimitMiddleware(rds, OtpRateLimitConfig()).Handler,
		NewAuthenticationMiddleware(m.co, p, di.Auth, di.Firebase).Handler,
		NewAuthorizationMiddleware(m.co).Handler,
	)
}

func NewMiddlewares(co shared.ContextOperator) Middlewares {
	return &mid{co}
}
