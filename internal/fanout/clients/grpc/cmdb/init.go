package integrationcmdb

import (
	config "github.com/borisbbtest/GoMon/internal/fanout/configs"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
	"github.com/borisbbtest/GoMon/internal/models/cmdb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewConnCmd Конструктор клиента к сервису CMDB
func NewConnCmd(cfg config.MainConfig) *ServiceWrapperCmdb {

	conn, err := grpc.Dial(cfg.ServiceGRpcCMD, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.Log.Fatal().Err(err).Msg("failed initialize connection")
	}

	return &ServiceWrapperCmdb{
		Ci: cmdb.NewCmdbClient(conn),
	}
}
