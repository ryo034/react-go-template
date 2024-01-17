package middleware

import (
	"github.com/ogen-go/ogen/middleware"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/config"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/logger"
	middl "github.com/ryo034/react-go-template/apps/system/api/infrastructure/middleware"
	"time"
)

type Middlewares interface {
	Global(conf config.Reader) []middleware.Middleware
}

type mid struct{}

func (m *mid) Global(conf config.Reader) []middleware.Middleware {
	logConf := logger.Config{TimeFormat: time.RFC3339, UTC: true}
	zl := logger.NewZeroLogger(logConf, conf.IsLocal(), conf.ServiceName())
	return []middleware.Middleware{
		RequestID(),
		Recovery(zl, conf.IsLocal()),
		Logger(zl),
		Cors(&middl.CORSInfo{AllowOrigins: conf.AllowOrigins()}),
		Language(conf.DefaultLanguage()),
	}
}

func NewMiddlewares() Middlewares {
	return &mid{}
}
