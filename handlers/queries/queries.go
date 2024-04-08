package queries

import (
	"github.com/cod3rboy/practice-cqrs/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type PatientQuery struct {
	// Patient Repository Here
	Logger *zap.Logger
}

type NewPatientQueryParams struct {
	fx.In
	Server *server.Server
	Logger *zap.Logger
	// Repository Dependency
}

func NewPatientQuery(params NewPatientQueryParams) *PatientQuery {
	query := &PatientQuery{Logger: params.Logger}
	params.Server.Router().GET("/patient/id/:id", query.ForID)
	params.Server.Router().GET("/patient/name/:name", query.ForName)
	params.Server.Router().GET("/patient/all", query.All)
	return query
}

func (q *PatientQuery) ForID(ctx *gin.Context) {
	q.Logger.Info("Invoked query PatientQuery.ForID")
}

func (q *PatientQuery) ForName(ctx *gin.Context) {
	q.Logger.Info("Invoked query PatientQuery.ForName")
}

func (q *PatientQuery) All(ctx *gin.Context) {
	q.Logger.Info("Invoked query PatientQuery.All")
}
