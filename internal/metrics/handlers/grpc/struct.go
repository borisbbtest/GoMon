package grpc

import (
	"github.com/borisbbtest/GoMon/internal/metrics/configs"
	"github.com/borisbbtest/GoMon/internal/metrics/storage"
	"github.com/borisbbtest/GoMon/internal/models/metrics"
)

type WrapperHandlerRPC struct {
	ServerConf *configs.MainConfig
	Storage    storage.Storage
	metrics.UnimplementedMetricsServer
}
