package models

import (
	"context"
	"fmt"
	"strings"

	"github.com/borisbbtest/GoMon/internal/fanin/service"
	pb "github.com/borisbbtest/GoMon/internal/models/mgrevent"
)

// PushEvent - метод, отправляющий event в cmdb с конвертацией Event в protobuf Event
func (cw *ConfigWrapper) PushEvent(ctx context.Context, event *Event) error {
	var req pb.PushRequest
	pbEvent := event.ToPB()
	req.Ev = pbEvent
	user := ctx.Value(FanInContextKey("login"))
	if user == nil {
		return service.ErrNoUserInContext
	}
	req.User = user.(string)
	_, err := cw.Conns.Events.Push(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

// PushBatchEvents - метод, отправляющий пакетно event в cmdb с конвертацией Event в protobuf Event
func (cw *ConfigWrapper) PushBatchEvents(ctx context.Context, events []Event) error {
	var req pb.PushBatchRequest
	var pbevents []*pb.Event
	for _, event := range events {
		pbevent := event.ToPB()
		pbevents = append(pbevents, pbevent)
	}
	req.Ev = pbevents
	user := ctx.Value(FanInContextKey("login"))
	if user == nil {
		return service.ErrNoUserInContext
	}
	req.User = user.(string)
	resp, err := cw.Conns.Events.PushBatch(ctx, &req)
	if err != nil {
		return err
	}
	// Так как gRPC возвращает либо response либо err, проверяется дополнительно внутренний код ответа приложения при пакетной загрузке
	if strings.ToLower(resp.Code) != "ok" {
		return fmt.Errorf("metrics return code: %s", resp.Code)
	}
	return nil
}
