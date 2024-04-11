package activities

import "go.uber.org/fx"

var ActivitiesModule = fx.Provide(
	NewPatientActivity,
)
