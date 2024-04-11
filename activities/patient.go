package activities

import (
	"context"
	"errors"
	"time"

	"github.com/cod3rboy/practice-cqrs/repositories"
	"github.com/cod3rboy/practice-cqrs/repositories/models"
	"github.com/google/uuid"
	"go.temporal.io/sdk/worker"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type PatientActivity struct {
	logger      *zap.Logger
	respository repositories.PatientRepository
}

type NewPatientActivityParams struct {
	fx.In
	Logger     *zap.Logger
	Repository repositories.PatientRepository
	Worker     worker.Worker
}

func NewPatientActivity(params NewPatientActivityParams) *PatientActivity {
	activity := &PatientActivity{
		logger:      params.Logger,
		respository: params.Repository,
	}

	params.Worker.RegisterActivity(activity.CreatePatientEvent)
	params.Worker.RegisterActivity(activity.TransferPatientEvent)
	params.Worker.RegisterActivity(activity.DischargePatientEvent)
	params.Worker.RegisterActivity(activity.UpdatePatientEvent)
	params.Worker.RegisterActivity(activity.ProjectPatient)

	return activity
}

type CreatePatientEventParams struct {
	ID   uuid.UUID
	Name string
	Ward int
	Age  int
}

func (a *PatientActivity) CreatePatientEvent(ctx context.Context, params CreatePatientEventParams) (string, error) {
	return a.respository.PublishCreatePatient(ctx, models.PatientEventData{
		Id:         params.ID.String(),
		Name:       params.Name,
		Ward:       &params.Ward,
		Age:        &params.Age,
		Discharged: new(bool),
	})
}

type TransferPatientEventParams struct {
	ID   uuid.UUID
	Ward int
}

func (a *PatientActivity) TransferPatientEvent(ctx context.Context, params TransferPatientEventParams) (string, error) {
	patient, err := a.respository.GetPatientByID(ctx, params.ID.String())
	if err != nil {
		a.logger.Error("error getting patient", zap.String("Patient ID", params.ID.String()), zap.Error(err))
		return "", errors.New("patient not found")
	}
	return a.respository.PublishPatientEvent(ctx, repositories.PublishPatientEventParams{
		PatientID:      patient.ID,
		CurrentVersion: uint64(patient.CurrentVersion),
		EventType:      models.TransferPatient,
		Data: models.PatientEventData{
			Id:   params.ID.String(),
			Ward: &params.Ward,
		},
	})
}

type DischargePatientEventParams struct {
	ID uuid.UUID
}

func (a *PatientActivity) DischargePatientEvent(ctx context.Context, params DischargePatientEventParams) (string, error) {
	patient, err := a.respository.GetPatientByID(ctx, params.ID.String())
	if err != nil {
		a.logger.Error("error getting patient", zap.String("Patient ID", params.ID.String()), zap.Error(err))
		return "", errors.New("patient not found")
	}

	discharged := true
	return a.respository.PublishPatientEvent(ctx, repositories.PublishPatientEventParams{
		PatientID:      patient.ID,
		CurrentVersion: uint64(patient.CurrentVersion),
		EventType:      models.DischargePatient,
		Data: models.PatientEventData{
			Id:         patient.ID,
			Discharged: &discharged,
		},
	})
}

type UpdatePatientEventParams struct {
	ID          uuid.UUID
	PatientName string
	PatientAge  *int
}

func (a *PatientActivity) UpdatePatientEvent(ctx context.Context, params UpdatePatientEventParams) (string, error) {
	patient, err := a.respository.GetPatientByID(ctx, params.ID.String())
	if err != nil {
		a.logger.Error("error getting patient", zap.String("Patient ID", params.ID.String()), zap.Error(err))
		return "", errors.New("patient not found")
	}

	return a.respository.PublishPatientEvent(ctx, repositories.PublishPatientEventParams{
		PatientID:      patient.ID,
		CurrentVersion: uint64(patient.CurrentVersion),
		EventType:      models.UpdatePatient,
		Data: models.PatientEventData{
			Id:   patient.ID,
			Name: params.PatientName,
			Age:  params.PatientAge,
		},
	})
}

func (a *PatientActivity) ProjectPatient(ctx context.Context, patientId uuid.UUID) error {
	events, err := a.respository.GetPatientEventsByID(ctx, patientId.String())
	if err != nil {
		a.logger.Error("error getting event stream for patient", zap.String("Stream ID", patientId.String()), zap.Error(err))
		return errors.New("error getting event stream for patient")
	}

	patient := models.Patient{
		ID: patientId.String(),
	}
	for _, event := range events {
		patient.CurrentVersion = int32(event.StreamPosition)
		event.Payload.ApplyPayloadTo(&patient)
	}
	patient.ProjectionTime = time.Now()

	patient, err = a.respository.WritePatient(ctx, patient)
	if err != nil {
		a.logger.Error("error saving patient projection", zap.String("Stream ID", patientId.String()), zap.Error(err))
		return errors.New("error saving patient projection")
	}

	return nil
}
