package middleware

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func newLogger(isLocal bool) *zap.Logger {
	logger, _ := zap.NewProduction()
	if isLocal {
		logger, _ = zap.NewDevelopment()
	}
	return logger
}

func Logger(isLocal bool) gin.HandlerFunc {
	logger := newLogger(isLocal)
	return func(c *gin.Context) {
		ginzap.Ginzap(logger, time.RFC3339, true)
		ginzap.RecoveryWithZap(logger, true)
	}
}
