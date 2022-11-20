package models

import (
	"fmt"

	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
	"github.com/jackc/pgx/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Обеъект для БД
type PGEvent struct {
	Id          pgtype.UUID
	Title       pgtype.Text
	Description pgtype.Text
	Source      pgtype.Text
	Status      int32
	Created     pgtype.Timestamp
	Update      pgtype.Timestamp
	Key         pgtype.Text
	KeyClose    pgtype.Text
	Assigned    pgtype.TextArray
	Severity    int32
	AutoRunner  pgtype.Text
	RelarionCi  pgtype.TextArray
}
type Events struct {
	EventsPG []*PGEvent
}

// Конвертируем данные из текста массива в БД в массив строк
func arrayTextToStringArray(v pgtype.TextArray) (res []string) {
	for _, v := range v.Elements {
		res = append(res, v.String)
	}
	return
}

// Конвертор тут который занимается подготавливает данные в GRPC
func (hook *Events) ConvertTogRpcEvent() (ev *[]*mgrevent.Event) {
	buff := []*mgrevent.Event{}
	for _, v := range hook.EventsPG {
		if v != nil {
			tmpeve := &mgrevent.Event{
				Uuid:        fmt.Sprintf("%x-%x-%x-%x-%x", v.Id.Bytes[0:4], v.Id.Bytes[4:6], v.Id.Bytes[6:8], v.Id.Bytes[8:10], v.Id.Bytes[10:16]),
				Title:       v.Title.String,
				Description: v.Description.String,
				Source:      v.Source.String,
				Status:      v.Status,
				Severity:    v.Severity,
				Created:     timestamppb.New(v.Created.Time),
				Update:      timestamppb.New(v.Update.Time),
				Key:         v.Key.String,
				KeyClose:    v.KeyClose.String,
				Assigned:    arrayTextToStringArray(v.Assigned),
				AutoRunner:  v.AutoRunner.String,
				RelarionCi:  arrayTextToStringArray(v.RelarionCi),
				CreatedBy:   "none",
				Count:       0,
			}
			buff = append(buff, tmpeve)
		}
	}
	ev = &buff
	return
}
