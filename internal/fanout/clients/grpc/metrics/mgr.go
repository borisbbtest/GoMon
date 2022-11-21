package integrationmetric

import (
	"context"
	"time"

	"github.com/borisbbtest/GoMon/internal/models/metrics"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GetMetricDuration - получаем меткри от сервиса в временном интервали
func (hook *ServiceWrapperMetric) GetMetricDuration(ctx context.Context, start time.Time, end time.Time) (resp *[]*Metric, err error) {
	req := metrics.GetBatchRequest{
		Start: timestamppb.New(start),
		End:   timestamppb.New(end),
	}

	metric_, err := hook.Metric.GetBatch(ctx, &req)
	if err != nil {
		return nil, err
	}
	order := []*Metric{}
	for _, v := range metric_.Item {
		buff := &Metric{
			Name:              v.Name,
			Value:             v.Value,
			Localtime:         v.Localtime.AsTime(),
			SourceTime:        v.SourceTime.AsTime(),
			SourceFromSystems: v.SourceFromSystems,
			RelationCi:        v.RelationCi,
			Uuid:              v.Uuid,
			Tp:                v.Tp.Enum().String(),
		}
		order = append(order, buff)
	}
	resp = &order
	return resp, nil
}
