package models

import (
	"fmt"

	"github.com/borisbbtest/GoMon/internal/models/metrics"
	"github.com/jackc/pgx/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Обеъект для БД
type PGMetric struct {
	Id         pgtype.UUID
	Name       pgtype.Text
	Value      pgtype.Bit
	Source     pgtype.Text
	LoadTime   pgtype.Timestamp
	SourceTime pgtype.Timestamp
	RelarionCi pgtype.Text
}
type Metrics struct {
	EventsPG []*PGMetric
}

// Конвертируем данные из текста массива в БД в массив строк
func arrayTextToStringArray(v pgtype.TextArray) (res []string) {
	for _, v := range v.Elements {
		res = append(res, v.String)
	}
	return
}

// Конвертор тут который занимается подготавливает данные в GRPC
func (hook *Metrics) ConvertTogRpcEvent() (ev *[]*metrics.Metric) {
	buff := []*metrics.Metric{}
	for _, v := range hook.EventsPG {
		if v != nil {
			tmpeve := &metrics.Metric{
				Uuid:              fmt.Sprintf("%x-%x-%x-%x-%x", v.Id.Bytes[0:4], v.Id.Bytes[4:6], v.Id.Bytes[6:8], v.Id.Bytes[8:10], v.Id.Bytes[10:16]),
				Name:              v.Name.String,
				SourceFromSystems: v.Source.String,
				Value:             v.Id.Bytes[:],
				Localtime:         timestamppb.New(v.LoadTime.Time),
				RelationCi:        v.RelarionCi.String,
				SourceTime:        timestamppb.New(v.SourceTime.Time),
			}
			buff = append(buff, tmpeve)
		}
	}
	ev = &buff
	return
}
