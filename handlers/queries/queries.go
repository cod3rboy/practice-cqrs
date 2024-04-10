package queries

import (
	"net/http"

	"github.com/cod3rboy/practice-cqrs/ent/patient"
	"github.com/cod3rboy/practice-cqrs/repositories"
	"github.com/cod3rboy/practice-cqrs/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type PatientQuery struct {
	Repository repositories.PatientRepository
	Logger     *zap.Logger
}

type NewPatientQueryParams struct {
	fx.In
	Server     *server.Server
	Logger     *zap.Logger
	Repository repositories.PatientRepository
}

func NewPatientQuery(params NewPatientQueryParams) *PatientQuery {
	query := &PatientQuery{Logger: params.Logger, Repository: params.Repository}
	params.Server.Router().GET("/patient/id/:id", query.ForID)
	params.Server.Router().GET("/patient/name/:name", query.ForName)
	params.Server.Router().GET("/patient/all", query.All)
	return query
}

func (q *PatientQuery) ForID(ctx *gin.Context) {
	q.Logger.Info("Invoked query PatientQuery.ForID")
	patientId := ctx.Param("id")
	patient, err := q.Repository.GetPatientByID(ctx, patientId)
	if err != nil {
		q.Logger.Info("failed to get patient", zap.String("patient id", patientId))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get patient",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"patient": patient,
	})
}

func (q *PatientQuery) ForName(ctx *gin.Context) {
	q.Logger.Info("Invoked query PatientQuery.ForName")
	patientName := ctx.Param("name")
	patient, err := q.Repository.GetPatientByPredicates(ctx, patient.NameEQ(patientName))
	if err != nil {
		q.Logger.Info("failed to get patient", zap.Any("patient name", patientName))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get patient",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"patient": patient,
	})
}

func (q *PatientQuery) All(ctx *gin.Context) {
	q.Logger.Info("Invoked query PatientQuery.All")
	patients, err := q.Repository.GetPatients(ctx)
	if err != nil {
		q.Logger.Error("failed to get all patients")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get all patients",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"patients": patients,
	})
}
