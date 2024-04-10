package models

import (
	"encoding/json"
	"time"

	"github.com/cod3rboy/practice-cqrs/eventstore"
)

type PatientEventType string

const (
	CreatePatient PatientEventType = "CREATE_PATIENT"
)

type PatientEventData struct {
	Id         string
	Name       string `json:"Name,omitempty"`
	Ward       *int   `json:"Ward,omitempty"`
	Age        *int   `json:"Age,omitempty"`
	Discharged *bool  `json:"Discharged,omitempty"`
}

func (data *PatientEventData) ToEvent(eventType PatientEventType) eventstore.Event {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return eventstore.Event{
		Type: string(eventType),
		Data: dataBytes,
	}
}

func (data *PatientEventData) FromEvent(event eventstore.Event) {
	err := json.Unmarshal(event.Data, data)
	if err != nil {
		panic(err)
	}
}

type Patient struct {
	ID             string
	Name           string
	WardNumber     int
	Age            int
	Discharged     bool
	CurrentVersion int32
	ProjectionTime time.Time
}

type PatientEvent struct {
	StreamPosition uint64
	EventType      PatientEventType
	CreatedDate    time.Time
	Payload        PatientEventPayload
}

type PatientEventPayload interface {
	WritePayloadFrom(PatientEventData)
}

func NewPatientEventPayload(eventType PatientEventType) PatientEventPayload {
	switch eventType {
	case CreatePatient:
		return &PatientCreatePayload{}
	}
	return nil
}

type PatientCreatePayload struct {
}

func (p *PatientCreatePayload) WritePayloadFrom(data PatientEventData) {}
