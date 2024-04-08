package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("ward"),
		field.Int("age"),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return nil
}
