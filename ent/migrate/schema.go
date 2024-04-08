// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "stream_id", Type: field.TypeString},
		{Name: "stream_position", Type: field.TypeUint64, Default: 0},
		{Name: "event_type", Type: field.TypeString},
		{Name: "data", Type: field.TypeBytes},
		{Name: "created_date", Type: field.TypeTime, Nullable: true, Default: "CURRENT_TIMESTAMP"},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "event_stream_id_stream_position",
				Unique:  true,
				Columns: []*schema.Column{EventsColumns[1], EventsColumns[2]},
			},
		},
	}
	// PatientsColumns holds the columns for the "patients" table.
	PatientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "ward", Type: field.TypeInt},
		{Name: "age", Type: field.TypeInt},
	}
	// PatientsTable holds the schema information for the "patients" table.
	PatientsTable = &schema.Table{
		Name:       "patients",
		Columns:    PatientsColumns,
		PrimaryKey: []*schema.Column{PatientsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EventsTable,
		PatientsTable,
	}
)

func init() {
}
