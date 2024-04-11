package app

import (
	"context"
	"fmt"
	"log"

	"github.com/cod3rboy/practice-cqrs/config"
	"go.temporal.io/sdk/client"
	logger "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type NewTemporalClientParams struct {
	fx.In
	Logger logger.Logger
	Config config.Config
}

func NewTemporalClient(params NewTemporalClientParams) client.Client {
	clientOptions := client.Options{
		HostPort:  fmt.Sprintf("%s:%d", params.Config.TemporalHost, params.Config.TemporalPort),
		Logger:    params.Logger,
		Namespace: params.Config.TemporalNamespace,
	}
	temporalClient, err := client.Dial(clientOptions)
	if err != nil {
		log.Fatalf("failed to create temporal client: %v", err)
	}
	return temporalClient
}

type TemporalLogger struct {
	*zap.Logger
}

func NewTemporalLogger(logger *zap.Logger) logger.Logger {
	return &TemporalLogger{
		Logger: logger.Named("Temporal.Client"),
	}
}

func (l *TemporalLogger) Debug(msg string, keyvals ...interface{}) {
	l.Logger.Debug(msg, l.zapFieldsFromKeyValues(keyvals)...)
}

func (l *TemporalLogger) Info(msg string, keyvals ...interface{}) {
	l.Logger.Info(msg, l.zapFieldsFromKeyValues(keyvals)...)
}

func (l *TemporalLogger) Warn(msg string, keyvals ...interface{}) {
	l.Logger.Warn(msg, l.zapFieldsFromKeyValues(keyvals)...)
}

func (l *TemporalLogger) Error(msg string, keyvals ...interface{}) {
	l.Logger.Error(msg, l.zapFieldsFromKeyValues(keyvals)...)
}

func (l *TemporalLogger) zapFieldsFromKeyValues(keyvalues []interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(keyvalues)/2)
	for i := 0; i < len(keyvalues); i += 2 {
		fields = append(fields, zap.Any(fmt.Sprintf("%v", keyvalues[i]), keyvalues[i+1]))
	}
	return fields
}

type NewWorkerParams struct {
	fx.In
	Temporal client.Client
	LC       fx.Lifecycle
}

func NewWorker(params NewWorkerParams) worker.Worker {
	w := worker.New(params.Temporal, "patients-worker", worker.Options{})

	params.LC.Append(fx.Hook{
		OnStart: func(context.Context) error {
			if err := w.Start(); err != nil {
				return err
			}
			return nil
		},
		OnStop: func(context.Context) error {
			w.Stop()
			return nil
		},
	})

	return w
}

var TemporalModule = fx.Provide(
	fx.Annotate(NewTemporalLogger, fx.As(new(logger.Logger))),
	NewTemporalClient,
	NewWorker,
)
