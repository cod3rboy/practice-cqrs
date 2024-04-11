package models

import (
	"encoding/json"
	"time"

	"github.com/cod3rboy/practice-cqrs/eventstore"
)

type PatientEventType string

const (
	CreatePatient    PatientEventType = "CREATE_PATIENT"
	TransferPatient  PatientEventType = "TRANSFER_PATIENT"
	DischargePatient PatientEventType = "DISCHARGE_PATIENT"
	UpdatePatient    PatientEventType = "UPDATE_PATIENT"
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
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	WardNumber     int       `json:"ward_number"`
	Age            int       `json:"age"`
	Discharged     bool      `json:"discharged"`
	CurrentVersion int32     `json:"current_version"`
	ProjectionTime time.Time `json:"projection_time"`
}

type PatientEvent struct {
	StreamPosition uint64
	EventType      PatientEventType
	CreatedDate    time.Time
	Payload        PatientEventPayload
}

type PatientEventPayload interface {
	WritePayloadFrom(PatientEventData)
	ApplyPayloadTo(*Patient)
}

func NewPatientEventPayload(eventType PatientEventType) PatientEventPayload {
	switch eventType {
	case CreatePatient:
		return &PatientCreatePayload{}
	case TransferPatient:
		return &PatientTransferPayload{}
	case DischargePatient:
		return &PatientDischargePayload{}
	case UpdatePatient:
		return &PatientUpdatePayload{}
	}
	return nil
}

type PatientCreatePayload struct {
	id   string
	name string
	ward int
	age  int
}

func (p *PatientCreatePayload) WritePayloadFrom(data PatientEventData) {
	p.id = data.Id
	p.name = data.Name
	p.ward = *data.Ward
	p.age = *data.Age
}

func (p *PatientCreatePayload) ApplyPayloadTo(patient *Patient) {
	patient.Name = p.name
	patient.WardNumber = p.ward
	patient.Age = p.age
}

type PatientTransferPayload struct {
	id   string
	ward int
}

func (p *PatientTransferPayload) WritePayloadFrom(data PatientEventData) {
	p.id = data.Id
	p.ward = *data.Ward
}

func (p *PatientTransferPayload) ApplyPayloadTo(patient *Patient) {
	patient.WardNumber = p.ward
}

type PatientDischargePayload struct {
	id string
}

func (p *PatientDischargePayload) WritePayloadFrom(data PatientEventData) {
	p.id = data.Id
}

func (p *PatientDischargePayload) ApplyPayloadTo(patient *Patient) {
	patient.Discharged = true
}

type PatientUpdatePayload struct {
	id   string
	name string
	age  *int
}

func (p *PatientUpdatePayload) WritePayloadFrom(data PatientEventData) {
	p.id = data.Id
	p.name = data.Name
	p.age = data.Age
}

func (p *PatientUpdatePayload) ApplyPayloadTo(patient *Patient) {
	if p.name != "" {
		patient.Name = p.name
	}

	if p.age != nil {
		patient.Age = *p.age
	}
}
