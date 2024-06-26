// Code generated by ent, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cod3rboy/practice-cqrs/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldID, id))
}

// StreamId applies equality check predicate on the "StreamId" field. It's identical to StreamIdEQ.
func StreamId(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldStreamId, v))
}

// StreamPosition applies equality check predicate on the "StreamPosition" field. It's identical to StreamPositionEQ.
func StreamPosition(v uint64) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldStreamPosition, v))
}

// EventType applies equality check predicate on the "EventType" field. It's identical to EventTypeEQ.
func EventType(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEventType, v))
}

// Data applies equality check predicate on the "Data" field. It's identical to DataEQ.
func Data(v []byte) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldData, v))
}

// CreatedDate applies equality check predicate on the "CreatedDate" field. It's identical to CreatedDateEQ.
func CreatedDate(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedDate, v))
}

// StreamIdEQ applies the EQ predicate on the "StreamId" field.
func StreamIdEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldStreamId, v))
}

// StreamIdNEQ applies the NEQ predicate on the "StreamId" field.
func StreamIdNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldStreamId, v))
}

// StreamIdIn applies the In predicate on the "StreamId" field.
func StreamIdIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldStreamId, vs...))
}

// StreamIdNotIn applies the NotIn predicate on the "StreamId" field.
func StreamIdNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldStreamId, vs...))
}

// StreamIdGT applies the GT predicate on the "StreamId" field.
func StreamIdGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldStreamId, v))
}

// StreamIdGTE applies the GTE predicate on the "StreamId" field.
func StreamIdGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldStreamId, v))
}

// StreamIdLT applies the LT predicate on the "StreamId" field.
func StreamIdLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldStreamId, v))
}

// StreamIdLTE applies the LTE predicate on the "StreamId" field.
func StreamIdLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldStreamId, v))
}

// StreamIdContains applies the Contains predicate on the "StreamId" field.
func StreamIdContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldStreamId, v))
}

// StreamIdHasPrefix applies the HasPrefix predicate on the "StreamId" field.
func StreamIdHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldStreamId, v))
}

// StreamIdHasSuffix applies the HasSuffix predicate on the "StreamId" field.
func StreamIdHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldStreamId, v))
}

// StreamIdEqualFold applies the EqualFold predicate on the "StreamId" field.
func StreamIdEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldStreamId, v))
}

// StreamIdContainsFold applies the ContainsFold predicate on the "StreamId" field.
func StreamIdContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldStreamId, v))
}

// StreamPositionEQ applies the EQ predicate on the "StreamPosition" field.
func StreamPositionEQ(v uint64) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldStreamPosition, v))
}

// StreamPositionNEQ applies the NEQ predicate on the "StreamPosition" field.
func StreamPositionNEQ(v uint64) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldStreamPosition, v))
}

// StreamPositionIn applies the In predicate on the "StreamPosition" field.
func StreamPositionIn(vs ...uint64) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldStreamPosition, vs...))
}

// StreamPositionNotIn applies the NotIn predicate on the "StreamPosition" field.
func StreamPositionNotIn(vs ...uint64) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldStreamPosition, vs...))
}

// StreamPositionGT applies the GT predicate on the "StreamPosition" field.
func StreamPositionGT(v uint64) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldStreamPosition, v))
}

// StreamPositionGTE applies the GTE predicate on the "StreamPosition" field.
func StreamPositionGTE(v uint64) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldStreamPosition, v))
}

// StreamPositionLT applies the LT predicate on the "StreamPosition" field.
func StreamPositionLT(v uint64) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldStreamPosition, v))
}

// StreamPositionLTE applies the LTE predicate on the "StreamPosition" field.
func StreamPositionLTE(v uint64) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldStreamPosition, v))
}

// EventTypeEQ applies the EQ predicate on the "EventType" field.
func EventTypeEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEventType, v))
}

// EventTypeNEQ applies the NEQ predicate on the "EventType" field.
func EventTypeNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldEventType, v))
}

// EventTypeIn applies the In predicate on the "EventType" field.
func EventTypeIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldEventType, vs...))
}

// EventTypeNotIn applies the NotIn predicate on the "EventType" field.
func EventTypeNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldEventType, vs...))
}

// EventTypeGT applies the GT predicate on the "EventType" field.
func EventTypeGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldEventType, v))
}

// EventTypeGTE applies the GTE predicate on the "EventType" field.
func EventTypeGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldEventType, v))
}

// EventTypeLT applies the LT predicate on the "EventType" field.
func EventTypeLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldEventType, v))
}

// EventTypeLTE applies the LTE predicate on the "EventType" field.
func EventTypeLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldEventType, v))
}

// EventTypeContains applies the Contains predicate on the "EventType" field.
func EventTypeContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldEventType, v))
}

// EventTypeHasPrefix applies the HasPrefix predicate on the "EventType" field.
func EventTypeHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldEventType, v))
}

// EventTypeHasSuffix applies the HasSuffix predicate on the "EventType" field.
func EventTypeHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldEventType, v))
}

// EventTypeEqualFold applies the EqualFold predicate on the "EventType" field.
func EventTypeEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldEventType, v))
}

// EventTypeContainsFold applies the ContainsFold predicate on the "EventType" field.
func EventTypeContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldEventType, v))
}

// DataEQ applies the EQ predicate on the "Data" field.
func DataEQ(v []byte) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "Data" field.
func DataNEQ(v []byte) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "Data" field.
func DataIn(vs ...[]byte) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "Data" field.
func DataNotIn(vs ...[]byte) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "Data" field.
func DataGT(v []byte) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "Data" field.
func DataGTE(v []byte) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "Data" field.
func DataLT(v []byte) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "Data" field.
func DataLTE(v []byte) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldData, v))
}

// CreatedDateEQ applies the EQ predicate on the "CreatedDate" field.
func CreatedDateEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedDate, v))
}

// CreatedDateNEQ applies the NEQ predicate on the "CreatedDate" field.
func CreatedDateNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldCreatedDate, v))
}

// CreatedDateIn applies the In predicate on the "CreatedDate" field.
func CreatedDateIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldCreatedDate, vs...))
}

// CreatedDateNotIn applies the NotIn predicate on the "CreatedDate" field.
func CreatedDateNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldCreatedDate, vs...))
}

// CreatedDateGT applies the GT predicate on the "CreatedDate" field.
func CreatedDateGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldCreatedDate, v))
}

// CreatedDateGTE applies the GTE predicate on the "CreatedDate" field.
func CreatedDateGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldCreatedDate, v))
}

// CreatedDateLT applies the LT predicate on the "CreatedDate" field.
func CreatedDateLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldCreatedDate, v))
}

// CreatedDateLTE applies the LTE predicate on the "CreatedDate" field.
func CreatedDateLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldCreatedDate, v))
}

// CreatedDateIsNil applies the IsNil predicate on the "CreatedDate" field.
func CreatedDateIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldCreatedDate))
}

// CreatedDateNotNil applies the NotNil predicate on the "CreatedDate" field.
func CreatedDateNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldCreatedDate))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(sql.NotPredicates(p))
}
