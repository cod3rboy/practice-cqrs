package handlers

import (
	"github.com/cod3rboy/practice-cqrs/handlers/commands"
	"github.com/cod3rboy/practice-cqrs/handlers/queries"
	"go.uber.org/fx"
)

var HandlersModule = fx.Provide(
	commands.NewPatientCommand,
	queries.NewPatientQuery,
)

type Commands struct {
	fx.In
	*commands.PatientCommand
}

type Queries struct {
	fx.In
	*queries.PatientQuery
}

var RegisterHandlers = fx.Invoke(func(Commands, Queries) {})
