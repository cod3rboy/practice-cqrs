package workflows

import (
	"time"

	"github.com/cod3rboy/practice-cqrs/activities"
	"github.com/google/uuid"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type PatientWorkflow struct {
	logger   *zap.Logger
	activity *activities.PatientActivity
}

type NewPatientWorkflowParams struct {
	fx.In
	Logger   *zap.Logger
	Activity *activities.PatientActivity
	Worker   worker.Worker
}

func NewPatientWorkflow(params NewPatientWorkflowParams) *PatientWorkflow {
	workflow := &PatientWorkflow{
		logger:   params.Logger,
		activity: params.Activity,
	}

	params.Worker.RegisterWorkflow(workflow.Admit)
	params.Worker.RegisterWorkflow(workflow.Transfer)
	params.Worker.RegisterWorkflow(workflow.Discharge)
	params.Worker.RegisterWorkflow(workflow.UpdateAge)
	params.Worker.RegisterWorkflow(workflow.UpdateName)
	params.Worker.RegisterWorkflow(workflow.Projection)

	return workflow
}

type PatientAdmitWorkflowParams struct {
	ID   uuid.UUID
	Name string
	Ward int
	Age  int
}

func (w *PatientWorkflow) Admit(ctx workflow.Context, params PatientAdmitWorkflowParams) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 3 * time.Second,
	}
	parentCtx := ctx
	ctx = workflow.WithActivityOptions(ctx, ao)

	var patientId string

	err := workflow.ExecuteActivity(ctx, w.activity.CreatePatientEvent, activities.CreatePatientEventParams{
		ID:   params.ID,
		Name: params.Name,
		Ward: params.Ward,
		Age:  params.Age,
	}).Get(ctx, &patientId)

	if err != nil {
		w.logger.Error("failed to execute activity", zap.Error(err))
		return err
	}

	err = workflow.ExecuteChildWorkflow(parentCtx, w.Projection, params.ID).Get(parentCtx, nil)
	if err != nil {
		w.logger.Error("parent execution received child execution failure.", zap.Error(err))
		return err
	}

	return nil
}

type PatientTransferWorkflowParams struct {
	ID   uuid.UUID
	Ward int
}

func (w *PatientWorkflow) Transfer(ctx workflow.Context, params PatientTransferWorkflowParams) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 3 * time.Second,
	}
	parentCtx := ctx
	ctx = workflow.WithActivityOptions(ctx, ao)

	var patientId string

	err := workflow.ExecuteActivity(ctx, w.activity.TransferPatientEvent, activities.TransferPatientEventParams{
		ID:   params.ID,
		Ward: params.Ward,
	}).Get(ctx, &patientId)

	if err != nil {
		w.logger.Error("failed to execute activity", zap.Error(err))
		return err
	}

	err = workflow.ExecuteChildWorkflow(parentCtx, w.Projection, params.ID).Get(parentCtx, nil)
	if err != nil {
		w.logger.Error("parent execution received child execution failure.", zap.Error(err))
		return err
	}

	return nil
}

type PatientDischargeWorkflowParams struct {
	ID uuid.UUID
}

func (w *PatientWorkflow) Discharge(ctx workflow.Context, params PatientDischargeWorkflowParams) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 3 * time.Second,
	}
	parentCtx := ctx
	ctx = workflow.WithActivityOptions(ctx, ao)

	var patientId string

	err := workflow.ExecuteActivity(ctx, w.activity.DischargePatientEvent, activities.DischargePatientEventParams{
		ID: params.ID,
	}).Get(ctx, &patientId)

	if err != nil {
		w.logger.Error("failed to execute activity", zap.Error(err))
		return err
	}

	err = workflow.ExecuteChildWorkflow(parentCtx, w.Projection, params.ID).Get(parentCtx, nil)
	if err != nil {
		w.logger.Error("parent execution received child execution failure.", zap.Error(err))
		return err
	}

	return nil
}

type PatientUpdateNameWorkflowParams struct {
	ID          uuid.UUID
	PatientName string
}

func (w *PatientWorkflow) UpdateName(ctx workflow.Context, params PatientUpdateNameWorkflowParams) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 3 * time.Second,
	}
	parentCtx := ctx
	ctx = workflow.WithActivityOptions(ctx, ao)

	var patientId string

	err := workflow.ExecuteActivity(ctx, w.activity.UpdatePatientEvent, activities.UpdatePatientEventParams{
		ID:          params.ID,
		PatientName: params.PatientName,
	}).Get(ctx, &patientId)

	if err != nil {
		w.logger.Error("failed to execute activity", zap.Error(err))
		return err
	}

	err = workflow.ExecuteChildWorkflow(parentCtx, w.Projection, params.ID).Get(parentCtx, nil)
	if err != nil {
		w.logger.Error("parent execution received child execution failure.", zap.Error(err))
		return err
	}

	return nil
}

type PatientUpdateAgeWorkflowParams struct {
	ID         uuid.UUID
	PatientAge int
}

func (w *PatientWorkflow) UpdateAge(ctx workflow.Context, params PatientUpdateAgeWorkflowParams) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 3 * time.Second,
	}
	parentCtx := ctx
	ctx = workflow.WithActivityOptions(ctx, ao)

	var patientId string

	err := workflow.ExecuteActivity(ctx, w.activity.UpdatePatientEvent, activities.UpdatePatientEventParams{
		ID:         params.ID,
		PatientAge: &params.PatientAge,
	}).Get(ctx, &patientId)

	if err != nil {
		w.logger.Error("failed to execute activity", zap.Error(err))
		return err
	}

	err = workflow.ExecuteChildWorkflow(parentCtx, w.Projection, params.ID).Get(parentCtx, nil)
	if err != nil {
		w.logger.Error("parent execution received child execution failure.", zap.Error(err))
		return err
	}

	return nil
}

func (w *PatientWorkflow) Projection(ctx workflow.Context, patientID uuid.UUID) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 3 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, w.activity.ProjectPatient, patientID).Get(ctx, nil)

	if err != nil {
		w.logger.Error("failed to execute activity", zap.Error(err))
		return err
	}

	return nil
}
