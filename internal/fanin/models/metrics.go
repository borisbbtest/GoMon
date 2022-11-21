package models

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/borisbbtest/GoMon/internal/models/metrics"
)

// PushBatchMetrics - метод, отправляющий metric в cmdb с конвертацией Metric в protobuf Metric
func (cw *ConfigWrapper) PushBatchMetrics(ctx context.Context, metrics []Metric) error {
	var req pb.PushBatchRequest
	var pbmetrics []*pb.Metric
	for _, metric := range metrics {
		pbmetric, err := metric.ToPB()
		if err != nil {
			return err
		}
		pbmetrics = append(pbmetrics, pbmetric)
	}
	req.Item = pbmetrics
	resp, err := cw.Conns.Metrics.PushBatch(ctx, &req)
	if err != nil {
		return err
	}
	// Так как gRPC возвращает либо response либо err, проверяется дополнительно внутренний код ответа приложения при пакетной загрузке
	if strings.ToLower(resp.Code) != "ok" {
		return fmt.Errorf("metrics return code: %s", resp.Code)
	}
	return nil
}
