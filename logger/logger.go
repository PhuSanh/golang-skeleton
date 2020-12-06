package logger

import (
	"go.uber.org/zap"
	"sync"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func Init() {
	once.Do(func() {
		logger, _ = zap.NewProduction()
		defer logger.Sync()
	})
}

func Client() *zap.Logger {
	return logger
}
