package commands

import (
	"github.com/cod3rboy/practice-cqrs/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type PatientCommand struct {
	// Patient Repository Here
	Logger *zap.Logger
}

type NewPatientCommandParams struct {
	fx.In
	Server *server.Server
	Logger *zap.Logger
	// Repository Dependency
}

func NewPatientCommand(params NewPatientCommandParams) *PatientCommand {
	cmd := &PatientCommand{Logger: params.Logger}
	params.Server.Router().POST("/patient/admit", cmd.Admit)
	params.Server.Router().POST("/patient/transfer", cmd.Transfer)
	params.Server.Router().POST("/patient/discharge", cmd.Discharge)
	params.Server.Router().POST("/patient/updateAge", cmd.UpdateAge)
	params.Server.Router().POST("/patient/updateName", cmd.UpdateName)
	return cmd
}

func (c *PatientCommand) Admit(ctx *gin.Context) {
	c.Logger.Info("Invoked command PatientCommand.Admit")
}

func (c *PatientCommand) Transfer(ctx *gin.Context) {
	c.Logger.Info("Invoked command PatientCommand.Transfer")
}

func (c *PatientCommand) Discharge(ctx *gin.Context) {
	c.Logger.Info("Invoked command PatientCommand.Discharge")
}

func (c *PatientCommand) UpdateAge(ctx *gin.Context) {
	c.Logger.Info("Invoked command PatientCommand.UpdateAge")
}

func (c *PatientCommand) UpdateName(ctx *gin.Context) {
	c.Logger.Info("Invoked command PatientCommand.UpdateName")
}
