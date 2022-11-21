package integrationmetric

import (
	"github.com/borisbbtest/GoMon/internal/models/metrics"
)

type ServiceWrapperMetric struct {
	Metric metrics.MetricsClient //соединений gRPC IDM
}
