package servergrpc

import (
	"net"

	"github.com/borisbbtest/GoMon/internal/mgrevent/configs"
	handler "github.com/borisbbtest/GoMon/internal/mgrevent/handlers/grpc"
	"github.com/borisbbtest/GoMon/internal/mgrevent/storage"
	"github.com/borisbbtest/GoMon/internal/mgrevent/utils"
	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
	"google.golang.org/grpc"
)

type serviceRPCEventMgr struct {
	wrapp handler.WrapperHandlerRPC
}

func NewRPC(cfg *configs.MainConfig, st storage.Storage) *serviceRPCEventMgr {
	return &serviceRPCEventMgr{
		wrapp: handler.WrapperHandlerRPC{
			ServerConf: cfg,
			Storage:    st,
		},
	}
}

func (hook *serviceRPCEventMgr) Start() (err error) {

	listen, err := net.Listen("tcp", hook.wrapp.ServerConf.RunAddress)
	if err != nil {
		utils.Log.Error().Err(err)
		return
	}
	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer()

	// регистрируем сервис
	mgrevent.RegisterEventsServer(s, &hook.wrapp)
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
