package app

import (
	"github.com/borisbbtest/GoMon/internal/fanin/configs"
	"github.com/borisbbtest/GoMon/internal/fanin/handlers"
	"github.com/borisbbtest/GoMon/internal/fanin/models"
	"github.com/borisbbtest/GoMon/internal/models/cmdb"
	"github.com/borisbbtest/GoMon/internal/models/events"
	"github.com/borisbbtest/GoMon/internal/models/idm"
	"github.com/borisbbtest/GoMon/internal/models/metrics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// createGRPCConnection - создает gRPC подключение на основе адреса приложения
func createGRPCConnection(adress string) *grpc.ClientConn {
	conn, err := grpc.Dial(adress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).Msg("failed initialize connection")
	}
	return conn
}

// initializeHTTP - создает подключения и клиенты gRPC для всех модулей приложения
func initializeHTTP(cfg *configs.AppConfig) (*handlers.HTTP, *ConnPool) {
	connpool := ConnPool{}
	clientpool := models.ClientPool{}
	if cfg.IdmAddress != "" {
		connpool.idm = createGRPCConnection(cfg.IdmAddress)
		clientpool.Idm = idm.NewIdmClient(connpool.idm)
	}
	if cfg.CmdbAddress != "" {
		connpool.cmdb = createGRPCConnection(cfg.CmdbAddress)
		clientpool.Cmdb = cmdb.NewCmdbClient(connpool.cmdb)
	}
	if cfg.MetricsAddress != "" {
		connpool.metrics = createGRPCConnection(cfg.MetricsAddress)
		clientpool.Metrics = metrics.NewMetricsClient(connpool.metrics)
	}
	if cfg.EventsAddress != "" {
		connpool.events = createGRPCConnection(cfg.EventsAddress)
		clientpool.Events = events.NewEventsClient(connpool.events)
	}
	h := handlers.HTTP{
		App: &models.ConfigWrapper{
			Cfg:   cfg,
			Conns: &clientpool,
		},
	}
	return &h, &connpool
}

// Close - закрывает все подключения gRPC если они были
func (p *ConnPool) Close() {
	if p.idm != nil {
		p.idm.Close()
	}
	if p.cmdb != nil {
		p.cmdb.Close()
	}
	if p.metrics != nil {
		p.metrics.Close()
	}
	if p.events != nil {
		p.events.Close()
	}
}
