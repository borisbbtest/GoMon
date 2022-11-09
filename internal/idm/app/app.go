package app

import (
	"context"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/borisbbtest/GoMon/internal/idm/configs"
	"github.com/borisbbtest/GoMon/internal/idm/database"
	"github.com/borisbbtest/GoMon/internal/idm/models"
	pb "github.com/borisbbtest/GoMon/internal/idm/proto/idm"
	"github.com/borisbbtest/GoMon/internal/idm/service"
)

var (
	buildVersion string = "N/A"
	buildDate    string = "N/A"
	buildCommit  string = "N/A"
	log                 = zerolog.New(service.LogConfig()).With().Timestamp().Caller().Logger()
)

func BuildApp() {
	log.Info().Msg("idm started")
	log.Info().Msg("Build version: " + buildVersion)
	log.Info().Msg("Build date: " + buildDate)
	log.Info().Msg("Build commit: " + buildCommit)
	cfg := configs.LoadAppConfig()
	log.Info().Dict("cfg", zerolog.Dict().
		Str("DBDSN", cfg.DBDSN).
		Str("ServerAddressGRPC", cfg.ServerAddressGRPC).
		Dur("SessionTimeExpired", cfg.SessionTimeExpired).
		Bool("ReInit", cfg.ReInit),
	).Msg("Server config")
	ctx := context.Background()
	repo, err := database.NewDBStorage(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed initialize db connection")
		return
	}
	cfgwrapper := &models.ConfigWrapper{
		Cfg:  cfg,
		Repo: repo,
	}
	grpcwrapper := GRPC{
		App: cfgwrapper,
	}
	listen, err := net.Listen("tcp", grpcwrapper.App.Cfg.ServerAddressGRPC)
	if err != nil {
		log.Fatal().Err(err).Msg("failed initialize gRPC listener")
	}
	srv := grpc.NewServer()
	pb.RegisterEventsServer(srv, &grpcwrapper)
	go func() {
		if err := srv.Serve(listen); err != nil && err != grpc.ErrServerStopped {
			log.Fatal().Err(err).Msg("failed initialize server")
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-sigChan
	wg := &sync.WaitGroup{}
	defer func() {
		grpcwrapper.App.Repo.Close()
	}()
	wg.Add(1)
	go func() {
		srv.GracefulStop()
		log.Info().Msg("grpc stopped")
		wg.Done()
	}()
	wg.Wait()
	log.Info().Msg("idm stopped")
}
