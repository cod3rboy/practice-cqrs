package repositories

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/cod3rboy/practice-cqrs/db"
	"github.com/cod3rboy/practice-cqrs/ent"
	"github.com/cod3rboy/practice-cqrs/ent/patient"
	"github.com/cod3rboy/practice-cqrs/ent/predicate"
	"github.com/cod3rboy/practice-cqrs/eventstore"
	"github.com/cod3rboy/practice-cqrs/repositories/models"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type PatientRepository interface {
	PublishCreatePatient(ctx context.Context, newPatient models.PatientEventData) (string, error)
	PublishPatientEvent(ctx context.Context, params PublishPatientEventParams) (string, error)
	WritePatient(ctx context.Context, patient models.Patient) (models.Patient, error)
	GetPatientByID(ctx context.Context, patientID string) (models.Patient, error)
	GetPatientByPredicates(ctx context.Context, predicates ...predicate.Patient) (models.Patient, error)
	GetPatientEventsByID(ctx context.Context, patientID string) ([]models.PatientEvent, error)
	GetPatients(ctx context.Context) ([]models.Patient, error)
}

type patientRepository struct {
	logger *zap.Logger
	es     eventstore.Store
	db     *db.DBClient
}

type NewPatientRepositoryParams struct {
	fx.In
	Logger     *zap.Logger
	EventStore eventstore.Store
	DBClient   *db.DBClient
}

func NewPatientRepository(params NewPatientRepositoryParams) PatientRepository {
	return &patientRepository{
		logger: params.Logger,
		es:     params.EventStore,
		db:     params.DBClient,
	}
}

func (r *patientRepository) PublishCreatePatient(ctx context.Context, newPatient models.PatientEventData) (string, error) {
	event := newPatient.ToEvent(models.CreatePatient)
	_, err := r.es.CreateStream(ctx, newPatient.Id, event)
	if err != nil {
		r.logger.Error("error creating stream event PublishCreatePatient", zap.Error(err))
		return "", err
	}

	return newPatient.Id, nil
}

type PublishPatientEventParams struct {
	PatientID      string
	CurrentVersion uint64
	EventType      models.PatientEventType
	Data           models.PatientEventData
}

func (r *patientRepository) PublishPatientEvent(ctx context.Context, params PublishPatientEventParams) (string, error) {
	event := params.Data.ToEvent(params.EventType)
	_, err := r.es.AppendStream(ctx, params.PatientID, params.CurrentVersion, event)
	if err != nil {
		r.logger.Error("error publishing event PublishPatientEvent", zap.Error(err))
		return "", err
	}
	return params.PatientID, nil
}

func (r *patientRepository) WritePatient(ctx context.Context, patient models.Patient) (result models.Patient, err error) {
	query := r.db.Patient.Create().
		SetID(uuid.MustParse(patient.ID)).
		SetName(patient.Name).
		SetWard(patient.WardNumber).
		SetAge(patient.Age).
		SetDischarged(patient.Discharged).
		SetCurrentVersion(patient.CurrentVersion).
		SetProjectedAtDatetime(patient.ProjectionTime).
		OnConflict(
			sql.ConflictColumns("id"),
			sql.ResolveWithNewValues(),
		).UpdateNewValues()

	id, _ := query.ID(ctx)

	projection, err := r.db.Patient.Get(ctx, id)
	if err != nil {
		r.logger.Error("error projecting patient", zap.String("Patient ID", patient.ID), zap.Error(err))
		return
	}

	result = PatientFromProjection(projection)

	return
}

func (r *patientRepository) GetPatientByID(ctx context.Context, patientID string) (result models.Patient, err error) {
	projection, err := r.db.Patient.Query().Where(patient.IDEQ(uuid.MustParse(patientID))).Only(ctx)

	if err != nil {
		r.logger.Warn("failed to find projection", zap.String("Patient ID", patientID), zap.Error(err))
		return
	}

	result = PatientFromProjection(projection)

	return
}

func (r *patientRepository) GetPatientByPredicates(ctx context.Context, predicates ...predicate.Patient) (result models.Patient, err error) {
	projection, err := r.db.Patient.Query().Where(predicates...).Only(ctx)

	if err != nil {
		r.logger.Warn("failed to find projection", zap.Any("predicates", predicates))
		return
	}

	result = PatientFromProjection(projection)

	return
}

func (r *patientRepository) GetPatientEventsByID(ctx context.Context, patientID string) ([]models.PatientEvent, error) {
	streamId := uuid.MustParse(patientID)
	events, err := r.es.ReadStreamEvents(ctx, streamId.String())
	if err != nil {
		r.logger.Error("failed to get stream events GetPatientEventsByID", zap.String("StreamID", patientID), zap.Error(err))
		return nil, err
	}

	patientEvents := make([]models.PatientEvent, 0, len(events))

	for _, event := range events {
		eventData := models.PatientEventData{}
		eventData.FromEvent(event.Event)
		eventPayload := models.NewPatientEventPayload(models.PatientEventType(event.Type))
		if eventPayload == nil {
			r.logger.Error("failed to create event payload GetPatientEventsByID", zap.String("StreamID", patientID), zap.Uint64("StreamPosition", event.StreamPosition))
			return nil, fmt.Errorf("failed to create event payload GetPatientEventsByID: StreamID=%s StreamPosition=%d", patientID, event.StreamPosition)
		}
		eventPayload.WritePayloadFrom(eventData)
		patientEvent := models.PatientEvent{
			StreamPosition: event.StreamPosition,
			EventType:      models.PatientEventType(event.Type),
			CreatedDate:    event.CreatedDate,
			Payload:        eventPayload,
		}
		patientEvents = append(patientEvents, patientEvent)
	}
	return patientEvents, nil
}

func (r *patientRepository) GetPatients(ctx context.Context) ([]models.Patient, error) {
	projections, err := r.db.Patient.Query().All(ctx)

	if err != nil {
		r.logger.Warn("failed to find projections", zap.Error(err))
		return nil, err
	}

	result := make([]models.Patient, 0, len(projections))

	for _, projection := range projections {
		result = append(result, PatientFromProjection(projection))
	}

	return result, nil
}

func PatientFromProjection(projection *ent.Patient) models.Patient {
	return models.Patient{
		ID:             projection.ID.String(),
		Name:           projection.Name,
		WardNumber:     projection.Ward,
		Age:            projection.Age,
		Discharged:     projection.Discharged,
		CurrentVersion: *projection.CurrentVersion,
		ProjectionTime: projection.ProjectedAtDatetime,
	}
}
