package eventstore

import (
	"context"
	"fmt"
	"time"

	"github.com/cod3rboy/practice-cqrs/db"
	"github.com/cod3rboy/practice-cqrs/ent"
	"github.com/cod3rboy/practice-cqrs/ent/event"
	"go.uber.org/fx"
)

type Store interface {
	CreateStream(ctx context.Context, streamID string, event Event) (Result, error)
	AppendStream(ctx context.Context, streamID string, currentVersion uint64, event Event) (Result, error)
	ReadStreamEvents(ctx context.Context, streamID string) ([]EventRecord, error)
}

type Event struct {
	Type string
	Data []byte
}

type EventRecord struct {
	StreamID       string
	StreamPosition uint64
	GlobalPosition uint64
	CreatedDate    time.Time
	Event
}

type Result struct {
	GlobalPosition uint64
	StreamPosition uint64
}

type pgStore struct {
	db *db.DBClient
}

type NewPostgresEventStoreParams struct {
	fx.In
	Db *db.DBClient
}

func NewPostgresEventStore(params NewPostgresEventStoreParams) Store {
	return &pgStore{
		db: params.Db,
	}
}

func (s *pgStore) CreateStream(ctx context.Context, streamID string, event Event) (r Result, e error) {
	builder := s.db.Event.Create().
		SetStreamId(streamID).
		SetStreamPosition(1).
		SetEventType(event.Type).
		SetData(event.Data)
	result, err := builder.Save(ctx)

	if err != nil {
		e = err
		return
	}

	if result.ID == 0 {
		e = fmt.Errorf("entity already exists: {StreamId: %s, StreamPosition: %d}", streamID, 1)
		return
	}

	r.GlobalPosition = uint64(result.ID)
	r.StreamPosition = result.StreamPosition

	return
}

func (s *pgStore) AppendStream(ctx context.Context, streamID string, currentVersion uint64, event Event) (r Result, e error) {
	nextStreamPosition := currentVersion + 1
	builder := s.db.Event.Create().
		SetStreamId(streamID).
		SetStreamPosition(nextStreamPosition).
		SetEventType(event.Type).
		SetData(event.Data)
	result, err := builder.Save(ctx)

	if err != nil {
		e = err
		return
	}

	if result.ID == 0 {
		e = fmt.Errorf("entity already exists: {StreamId: %s, StreamPosition: %d}", streamID, nextStreamPosition)
		return
	}

	r.GlobalPosition = uint64(result.ID)
	r.StreamPosition = result.StreamPosition

	return
}

func (s *pgStore) ReadStreamEvents(ctx context.Context, streamID string) ([]EventRecord, error) {
	events, err := s.db.Event.Query().
		Where(event.StreamIdEQ(streamID)).
		Order(ent.Asc(event.FieldID)). // oldest to latest order
		All(ctx)

	if err != nil {
		return nil, err
	}

	records := make([]EventRecord, 0, len(events))

	for _, e := range events {
		records = append(records, EventRecord{
			StreamID:       e.StreamId,
			StreamPosition: e.StreamPosition,
			GlobalPosition: uint64(e.ID),
			CreatedDate:    e.CreatedDate,
			Event: Event{
				Type: e.EventType,
				Data: e.Data,
			},
		})
	}

	return records, nil
}
