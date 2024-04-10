package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Immutable(),
		field.String("name"),
		field.Int("ward"),
		field.Int("age"),
		field.Bool("discharged"),
		field.Int32("current_version").Optional().Nillable(),
		field.Time("projected_at_datetime"),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return nil
}
