package app

import (
	"github.com/cod3rboy/practice-cqrs/config"
	"go.uber.org/zap"
)

func NewLogger(config config.Config) *zap.Logger {
	if config.Environment == "dev" {
		return zap.NewExample()
	}
	logger, _ := zap.NewProduction()
	return logger
}
