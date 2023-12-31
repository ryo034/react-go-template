package logger

import (
	"go.uber.org/zap"
)

func NewZapLogger() (*zap.Logger, error) {
	//TODO: 綺麗なログを出力できるようにする
	return zap.NewProduction()
}
