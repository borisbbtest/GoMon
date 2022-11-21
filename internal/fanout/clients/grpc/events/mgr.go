// Package пакет по работе с сервисом event manager
package integrationevents

import (
	"context"
	"time"

	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func pbEventMapEvent(v *mgrevent.Event) *Event {
	return &Event{
		Title:       v.Title,
		Description: v.Description,
		Source:      v.Source,
		Status:      v.Status,
		Created:     v.Created.AsTime(),
		Update:      v.Update.AsTime(),
		Key:         v.Key,
		KeyClose:    v.KeyClose,
		Assigned:    v.Assigned,
		AutoRunner:  v.AutoRunner,
		Severity:    v.Severity,
		RelarionCi:  v.RelarionCi,
		CreatedBy:   v.CreatedBy,
		Count:       v.Count,
		Uuid:        v.Uuid,
	}
}

// GetEventDuration - получаем меткри от сервиса во временном интервали
func (hook *ServiceWrapperEvents) GetEventsDuration(ctx context.Context, start time.Time, end time.Time) (resp *[]*Event, err error) {
	req := mgrevent.GetBatchRequest{
		Start: timestamppb.New(start),
		End:   timestamppb.New(end),
	}

	events_, err := hook.Events.GetBatch(ctx, &req)
	if err != nil {
		return nil, err
	}
	order := []*Event{}
	for _, v := range events_.Ev {
		buff := pbEventMapEvent(v)
		order = append(order, buff)
	}
	resp = &order
	return resp, nil
}

// GetEvent - получает одни евен по uuid
func (hook *ServiceWrapperEvents) GetEvent(ctx context.Context, uuid string) (resp *Event, err error) {
	req := mgrevent.GetRequest{
		Id: uuid,
	}
	v, err := hook.Events.Get(ctx, &req)
	if err != nil {
		return nil, err
	}
	return pbEventMapEvent(v.Ev), nil
}
