package integrationevents

import (
	config "github.com/borisbbtest/GoMon/internal/fanout/configs"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewConnIdm(cfg config.MainConfig) *ServiceWrapperEvents {

	conn, err := grpc.Dial(cfg.ServiceGRpcCMD, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.Log.Fatal().Err(err).Msg("failed initialize connection")
	}

	return &ServiceWrapperEvents{
		Events: mgrevent.NewEventsClient(conn),
	}
}
