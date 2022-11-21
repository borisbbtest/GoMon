package storage

import (
	"context"
	"time"

	"github.com/borisbbtest/GoMon/internal/metrics/models"
	"github.com/borisbbtest/GoMon/internal/models/metrics"
)

// Интерфейс стореджа
type Storage interface {
	Close()
	SaveEvents(ctx context.Context, eve []*metrics.Metric) (err error, qerr error)
	GetMetricsDuration(ctx context.Context, startTime time.Time, endTime time.Time) (err error, res *models.Metrics)
}
