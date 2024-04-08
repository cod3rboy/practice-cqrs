// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/cod3rboy/practice-cqrs/ent/event"
	"github.com/cod3rboy/practice-cqrs/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescStreamId is the schema descriptor for StreamId field.
	eventDescStreamId := eventFields[0].Descriptor()
	// event.StreamIdValidator is a validator for the "StreamId" field. It is called by the builders before save.
	event.StreamIdValidator = eventDescStreamId.Validators[0].(func(string) error)
	// eventDescStreamPosition is the schema descriptor for StreamPosition field.
	eventDescStreamPosition := eventFields[1].Descriptor()
	// event.DefaultStreamPosition holds the default value on creation for the StreamPosition field.
	event.DefaultStreamPosition = eventDescStreamPosition.Default.(uint64)
	// eventDescEventType is the schema descriptor for EventType field.
	eventDescEventType := eventFields[2].Descriptor()
	// event.EventTypeValidator is a validator for the "EventType" field. It is called by the builders before save.
	event.EventTypeValidator = eventDescEventType.Validators[0].(func(string) error)
}
