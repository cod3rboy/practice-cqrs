package commands

import (
	"context"
	"net/http"

	"github.com/cod3rboy/practice-cqrs/server"
	"github.com/cod3rboy/practice-cqrs/workflows"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type PatientCommand struct {
	Temporal  client.Client
	Workflow  *workflows.PatientWorkflow
	Logger    *zap.Logger
	TaskQueue string
}

type NewPatientCommandParams struct {
	fx.In
	Server   *server.Server
	Logger   *zap.Logger
	Temporal client.Client
	Workflow *workflows.PatientWorkflow
}

func NewPatientCommand(params NewPatientCommandParams) *PatientCommand {
	cmd := &PatientCommand{
		Logger:    params.Logger,
		Temporal:  params.Temporal,
		Workflow:  params.Workflow,
		TaskQueue: "patients-worker",
	}
	params.Server.Router().POST("/patient/admit", cmd.Admit)
	params.Server.Router().POST("/patient/transfer", cmd.Transfer)
	params.Server.Router().POST("/patient/discharge", cmd.Discharge)
	params.Server.Router().POST("/patient/updateAge", cmd.UpdateAge)
	params.Server.Router().POST("/patient/updateName", cmd.UpdateName)
	return cmd
}

func (c *PatientCommand) Admit(ctx *gin.Context) {
	c.Logger.Info("Invoked command PatientCommand.Admit")
	payload := struct {
		Name string `json:"name" binding:"required"`
		Ward int    `json:"ward" binding:"required"`
		Age  int    `json:"age" binding:"required"`
	}{}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	patientId := uuid.New()

	wo := client.StartWorkflowOptions{
		ID:        "AdmitPatient_" + patientId.String(),
		TaskQueue: c.TaskQueue,
	}

	c.Logger.Info("starting workflow", zap.String("workflow id", wo.ID))

	_, err := c.Temporal.ExecuteWorkflow(
		context.Background(),
		wo,
		c.Workflow.Admit,
		workflows.PatientAdmitWorkflowParams{
			ID:   patientId,
			Name: payload.Name,
			Ward: payload.Ward,
			Age:  payload.Age,
		},
	)
	if err != nil {
		c.Logger.Error("failed to start workflow", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": patientId.String(),
	})
}

func (c *PatientCommand) Transfer(ctx *gin.Context) {
	c.Logger.Info("Invoked command PatientCommand.Transfer")
	payload := struct {
		ID   string `json:"id" binding:"required"`
		Ward int    `json:"ward" binding:"required"`
	}{}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	patientId, idErr := uuid.Parse(payload.ID)
	if idErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid patient id",
		})
		return
	}

	wo := client.StartWorkflowOptions{
		ID:        "TransferPatient_" + uuid.NewString(),
		TaskQueue: c.TaskQueue,
	}

	c.Logger.Info("starting workflow", zap.String("workflow id", wo.ID))

	_, err := c.Temporal.ExecuteWorkflow(
		context.Background(),
		wo,
		c.Workflow.Transfer,
		workflows.PatientTransferWorkflowParams{
			ID:   patientId,
			Ward: payload.Ward,
		},
	)

	if err != nil {
		c.Logger.Error("failed to start workflow", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": patientId.String(),
	})
}

func (c *PatientCommand) Discharge(ctx *gin.Context) {
	c.Logger.Info("Invoked command PatientCommand.Discharge")
	payload := struct {
		ID string `json:"id" binding:"required"`
	}{}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	patientId, idErr := uuid.Parse(payload.ID)
	if idErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid patient id",
		})
		return
	}

	wo := client.StartWorkflowOptions{
		ID:        "DischargePatient_" + uuid.NewString(),
		TaskQueue: c.TaskQueue,
	}

	c.Logger.Info("starting workflow", zap.String("workflow id", wo.ID))

	_, err := c.Temporal.ExecuteWorkflow(
		context.Background(),
		wo,
		c.Workflow.Discharge,
		workflows.PatientDischargeWorkflowParams{
			ID: patientId,
		},
	)

	if err != nil {
		c.Logger.Error("failed to start workflow", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": patientId.String(),
	})
}

func (c *PatientCommand) UpdateAge(ctx *gin.Context) {
	c.Logger.Info("Invoked command PatientCommand.UpdateAge")
	payload := struct {
		ID  string `json:"id" binding:"required"`
		Age int    `json:"age" binding:"required"`
	}{}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	patientId, idErr := uuid.Parse(payload.ID)
	if idErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid patient id",
		})
		return
	}

	wo := client.StartWorkflowOptions{
		ID:        "PatientUpdateAge_" + uuid.NewString(),
		TaskQueue: c.TaskQueue,
	}

	c.Logger.Info("starting workflow", zap.String("workflow id", wo.ID))

	_, err := c.Temporal.ExecuteWorkflow(
		context.Background(),
		wo,
		c.Workflow.UpdateAge,
		workflows.PatientUpdateAgeWorkflowParams{
			ID:         patientId,
			PatientAge: payload.Age,
		},
	)

	if err != nil {
		c.Logger.Error("failed to start workflow", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": patientId.String(),
	})
}

func (c *PatientCommand) UpdateName(ctx *gin.Context) {
	c.Logger.Info("Invoked command PatientCommand.UpdateName")
	payload := struct {
		ID   string `json:"id" binding:"required"`
		Name string `json:"name" binding:"required"`
	}{}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	patientId, idErr := uuid.Parse(payload.ID)
	if idErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid patient id",
		})
		return
	}

	wo := client.StartWorkflowOptions{
		ID:        "PatientUpdateName_" + uuid.NewString(),
		TaskQueue: c.TaskQueue,
	}

	c.Logger.Info("starting workflow", zap.String("workflow id", wo.ID))

	_, err := c.Temporal.ExecuteWorkflow(
		context.Background(),
		wo,
		c.Workflow.UpdateName,
		workflows.PatientUpdateNameWorkflowParams{
			ID:          patientId,
			PatientName: payload.Name,
		},
	)

	if err != nil {
		c.Logger.Error("failed to start workflow", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": patientId.String(),
	})
}
