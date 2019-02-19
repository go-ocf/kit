package test

import (
	"context"
	"errors"

	"github.com/go-ocf/cqrs/event"
	"github.com/go-ocf/cqrs/eventstore"
)

type MockEventStore struct {
	events map[string]map[string][]event.EventUnmarshaler
}

func (s *MockEventStore) Save(ctx context.Context, groupId, aggregateId string, events []event.Event) (concurrencyException bool, err error) {
	return false, errors.New("not supported")
}

func (s *MockEventStore) SaveSnapshot(ctx context.Context, groupId, aggregateId string, event event.Event) (concurrencyException bool, err error) {
	return false, errors.New("not supported")
}

func (s *MockEventStore) LoadFromVersion(ctx context.Context, queries []eventstore.QueryFromVersion, eventHandler event.Handler) error {
	aggregates := make(map[string][]event.EventUnmarshaler)
	for _, device := range s.events {
		for aggrId, events := range device {
			aggregates[aggrId] = events
		}
	}

	var events []event.EventUnmarshaler
	for _, q := range queries {
		var ok bool
		var r []event.EventUnmarshaler

		if r, ok = aggregates[q.AggregateId]; !ok {
			continue
		}

		events = append(events, r...)
	}

	return eventHandler.Handle(ctx, &iter{events: events})
}

func makeModelId(groupId, aggregateId string) string {
	return groupId + "." + aggregateId
}

func (s *MockEventStore) allModels(queriesInt map[string]eventstore.QueryFromVersion) map[string]eventstore.QueryFromVersion {
	for groupId, group := range s.events {
		for aggrId, events := range group {
			queriesInt[makeModelId(groupId, aggrId)] = eventstore.QueryFromVersion{AggregateId: aggrId, Version: events[0].Version}
		}
	}
	return queriesInt
}

func (s *MockEventStore) LoadFromSnapshot(ctx context.Context, queries []eventstore.QueryFromSnapshot, eventHandler event.Handler) error {
	queriesInt := make(map[string]eventstore.QueryFromVersion)
	if len(queries) == 0 {
		queriesInt = s.allModels(queriesInt)
	} else {
		for _, query := range queries {
			switch {
			case query.GroupId == "" && query.AggregateId == "":
				queriesInt = s.allModels(queriesInt)
				break
			case query.GroupId != "" && query.AggregateId == "":
				if aggregates, ok := s.events[query.GroupId]; ok {
					for aggrId, events := range aggregates {
						queriesInt[makeModelId(query.GroupId, aggrId)] = eventstore.QueryFromVersion{AggregateId: aggrId, Version: events[0].Version}
					}
				}
			default:
				if aggregates, ok := s.events[query.GroupId]; ok {
					if events, ok := aggregates[query.AggregateId]; ok {
						queriesInt[makeModelId(query.GroupId, query.AggregateId)] = eventstore.QueryFromVersion{AggregateId: query.AggregateId, Version: events[0].Version}
					}
				}
			}
		}
	}

	ret := make([]eventstore.QueryFromVersion, 0, len(queriesInt))
	for _, q := range queriesInt {
		ret = append(ret, q)
	}

	return s.LoadFromVersion(ctx, ret, eventHandler)
}

type iter struct {
	idx    int
	events []event.EventUnmarshaler
}

func (i *iter) Next(ctx context.Context, eu *event.EventUnmarshaler) bool {
	if i.idx >= len(i.events) {
		return false
	}
	*eu = i.events[i.idx]
	i.idx++
	return true
}

func (i *iter) Err() error {
	return nil
}

func (s *MockEventStore) GetInstanceId(ctx context.Context, deviceId, resourceId string) (int64, error) {
	return -1, errors.New("not supported")
}
func (s *MockEventStore) RemoveInstanceId(ctx context.Context, instanceId int64) error {
	return errors.New("not supported")
}

func NewMockEventStore() *MockEventStore {
	return &MockEventStore{make(map[string]map[string][]event.EventUnmarshaler)}
}

func (e *MockEventStore) Append(deviceId, resourceId string, ev event.EventUnmarshaler) {
	var m map[string][]event.EventUnmarshaler
	var ok bool
	if m, ok = e.events[deviceId]; !ok {
		m = make(map[string][]event.EventUnmarshaler)
		e.events[deviceId] = m
	}
	var r []event.EventUnmarshaler
	if r, ok = m[resourceId]; !ok {
		r = make([]event.EventUnmarshaler, 0, 10)
	}
	m[resourceId] = append(r, ev)
}