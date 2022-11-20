package servergrpc

import (
	"net"

	"github.com/borisbbtest/GoMon/internal/metrics/configs"
	handler "github.com/borisbbtest/GoMon/internal/metrics/handlers/grpc"
	"github.com/borisbbtest/GoMon/internal/metrics/storage"
	"github.com/borisbbtest/GoMon/internal/metrics/utils"
	"github.com/borisbbtest/GoMon/internal/models/metrics"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
)

type serviceRPCMetricMgr struct {
	wrapp handler.WrapperHandlerRPC
}

func NewRPC(cfg *configs.MainConfig, st storage.Storage) *serviceRPCMetricMgr {
	return &serviceRPCMetricMgr{
		wrapp: handler.WrapperHandlerRPC{
			ServerConf: cfg,
			Storage:    st,
		},
	}
}

func (hook *serviceRPCMetricMgr) Start() (err error) {

	listen, err := net.Listen("tcp", hook.wrapp.ServerConf.RunAddress)
	if err != nil {
		utils.Log.Error().Err(err)
		return
	}
	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			//grpc_ctxtags.StreamServerInterceptor(),
			// grpc_opentracing.StreamServerInterceptor(),
			// grpc_prometheus.StreamServerInterceptor,
			// grpc_zap.StreamServerInterceptor(zapLogger),
			// grpc_auth.StreamServerInterceptor(myAuthFunction),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			//grpc_ctxtags.UnaryServerInterceptor(),
			// grpc_opentracing.UnaryServerInterceptor(),
			// grpc_prometheus.UnaryServerInterceptor,
			// grpc_zap.UnaryServerInterceptor(zapLogger),
			// grpc_auth.UnaryServerInterceptor(myAuthFunction),
			grpc_recovery.UnaryServerInterceptor(),
		)),
		grpc.ChainUnaryInterceptor(logging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(utils.Log))),
	)

	// регистрируем сервис
	metrics.RegisterMetricsServer(s, &hook.wrapp)
	utils.Log.Info().Msgf("Server gRPC is running ")

	// получаем запрос gRPC
	err = s.Serve(listen)
	if err != nil {
		utils.Log.Error().Err(err)
		return
	}
	utils.Log.Info().Msgf("End GRPC")
	defer s.Stop()
	defer listen.Close()
	defer hook.wrapp.Storage.Close()
	return
}
