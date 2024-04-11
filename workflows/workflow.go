package workflows

import "go.uber.org/fx"

var WorkflowsModule = fx.Provide(
	NewPatientWorkflow,
)
