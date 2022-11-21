package integrationidm

import (
	config "github.com/borisbbtest/GoMon/internal/fanout/configs"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
	"github.com/borisbbtest/GoMon/internal/models/idm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewConnIdm(cfg config.MainConfig) *ServiceWrapperIdm {

	conn, err := grpc.Dial(cfg.ServiceGRpcIDM, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.Log.Fatal().Err(err).Msg("failed initialize connection")
	}

	return &ServiceWrapperIdm{
		Idm: idm.NewIdmClient(conn),
	}
}
