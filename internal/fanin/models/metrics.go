package models

import (
	"context"
	"fmt"

	pb "github.com/borisbbtest/GoMon/internal/models/metrics"
)

// PushMetrics - метод, отправляющий metric в cmdb с конвертацией Metric в protobuf Metric
func (cw *ConfigWrapper) PushMetrics(ctx context.Context, metric *Metric) error {
	var req pb.PushRequest
	req.Item = metric.ToPB()
	_, err := cw.Conns.Metrics.Push(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

// PushBatchMetrics - метод, отправляющий metric в cmdb с конвертацией Metric в protobuf Metric
func (cw *ConfigWrapper) PushBatchMetrics(ctx context.Context, metrics []Metric) error {
	var req pb.PushBatchRequest
	var pbmetrics []*pb.Metric
	for _, metric := range metrics {
		pbmetric := metric.ToPB()
		pbmetrics = append(pbmetrics, pbmetric)
	}
	req.Item = pbmetrics
	resp, err := cw.Conns.Metrics.PushBatch(ctx, &req)
	if err != nil {
		return err
	}
	// Так как gRPC возвращает либо response либо err, проверяется дополнительно внутренний код ответа приложения при пакетной загрузке
	if resp.Code != "OK" {
		return fmt.Errorf("metrics return code: %s", resp.Code)
	}
	return nil
}
