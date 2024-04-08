package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("StreamId").NotEmpty().Immutable(),
		field.Uint64("StreamPosition").Default(0).Immutable(),
		field.String("EventType").NotEmpty().Immutable(),
		field.Bytes("Data").Immutable(),
		field.Time("CreatedDate").Annotations(
			&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			},
		).Optional().Immutable(),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return nil
}

// Indexes of the Event.
func (Event) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("StreamId", "StreamPosition").Unique(),
	}
}
