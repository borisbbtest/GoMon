package server_grpc

import (
	"net"

	"github.com/borisbbtest/GoMon/internal/metrics/utils"
	"github.com/borisbbtest/go_home_work/internal/config"
	handlersgrpc "github.com/borisbbtest/go_home_work/internal/handlers/handlers_grpc"
	"github.com/borisbbtest/go_home_work/internal/proto/shortrpc"
	"github.com/borisbbtest/go_home_work/internal/storage"
	"google.golang.org/grpc"
)

type serviceRPCShortURL struct {
	wrapp handlersgrpc.WrapperHandlerRPC
}

func NewRPC(cfg *config.ServiceShortURLConfig, st storage.Storage) *serviceRPCShortURL {
	return &serviceRPCShortURL{
		wrapp: handlersgrpc.WrapperHandlerRPC{
			ServerConf: cfg,
			Storage:    st,
		},
	}
}

func (hook *serviceRPCShortURL) Start() (err error) {

	listen, err := net.Listen("tcp", hook.wrapp.ServerConf.ServerRPC)
	if err != nil {
		utils.Log.Error().Err(err)
		return
	}
	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer()

	// регистрируем сервис
	shortrpc.RegisterShortURLServer(s, &hook.wrapp)

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
