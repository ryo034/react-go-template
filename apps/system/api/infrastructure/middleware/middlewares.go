package middleware

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/logger"
	"net/http"
)

type Middlewares interface {
	Global(h http.Handler, conf config.Reader, zl logger.Logger) http.Handler
}

type HttpMiddleware func(http.Handler) http.Handler

type mid struct{}

func applyMiddlewares(h http.Handler, middlewares ...HttpMiddleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func (m *mid) Global(h http.Handler, conf config.Reader, zl logger.Logger) http.Handler {
	return applyMiddlewares(
		h,
		NewRequestIDMiddleware().Handler,
		NewLangMiddleware(conf.DefaultLanguage()).Handler,
		NewLogMiddleware(zl, conf.IsLocal()).Handler,
		conf.Cors().Handler,
	)
}

func NewMiddlewares() Middlewares {
	return &mid{}
}
