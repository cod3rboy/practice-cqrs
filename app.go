package app

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/cod3rboy/practice-cqrs/activities"
	"github.com/cod3rboy/practice-cqrs/config"
	"github.com/cod3rboy/practice-cqrs/db"
	"github.com/cod3rboy/practice-cqrs/eventstore"
	"github.com/cod3rboy/practice-cqrs/handlers"
	"github.com/cod3rboy/practice-cqrs/repositories"
	"github.com/cod3rboy/practice-cqrs/server"
	"github.com/cod3rboy/practice-cqrs/workflows"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func SetupAndRun() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			NewLogger,
			db.NewDatabaseClient,
			server.NewServer,
			fx.Annotate(eventstore.NewPostgresEventStore, fx.As(new(eventstore.Store))),
			fx.Annotate(repositories.NewPatientRepository, fx.As(new(repositories.PatientRepository))),
		),
		TemporalModule,
		activities.ActivitiesModule,
		workflows.WorkflowsModule,
		handlers.HandlersModule,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Invoke(func(*server.Server) {}),
		handlers.RegisterHandlers,
	)

	app.Run()
}
