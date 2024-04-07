package app

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/cod3rboy/practice-cqrs/config"
	"github.com/cod3rboy/practice-cqrs/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func SetupAndRun() {
	app := fx.New(
		fx.Provide(config.NewConfig, NewLogger, server.NewServer),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Invoke(func(*server.Server) {}),
	)

	app.Run()
}

func NewLogger(config config.Config) *zap.Logger {
	if config.Environment == "dev" {
		return zap.NewExample()
	}
	logger, _ := zap.NewProduction()
	return logger
}
