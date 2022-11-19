package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/borisbbtest/GoMon/internal/fanin/configs"
	"github.com/borisbbtest/GoMon/internal/fanin/service"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

var (
	buildVersion string = "N/A"
	buildDate    string = "N/A"
	buildCommit  string = "N/A"
	log                 = zerolog.New(service.LogConfig()).With().Timestamp().Caller().Logger()
)

// HTTPServer - формирует роутер для HTTP и возвращает сервер и пул соединений для того, что бы их можно было закрыть.
func HTTPServer(ctx context.Context, cfg *configs.AppConfig) (*http.Server, *ConnPool) {
	h, connpool := initializeHTTP(cfg)
	r := chi.NewRouter()
	r.Use(h.GzipHandle)
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Get("/", h.HelloHandler)
	r.Post("/api/register", h.RegisterHandler)
	r.Post("/api/authorize", h.AuthorizeHandler)
	r.Route("/api", func(r chi.Router) {
		r.Use(h.CheckAuthorized)
		r.Post("/ci/push", h.PushCiHandler)
		r.Post("/ci/push/batch", h.PushBatchCiHandler)
		r.Post("/ci/delete", h.DeleteCiHandler)
		r.Post("/ci/delete/batch", h.DeleteBatchCiHandler)
		r.Post("/metric/push", h.PushMetricHandler)
		r.Post("/metric/push/batch", h.PushBatchMetricHandler)
		r.Post("/event/push", h.PushEventHandler)
		r.Post("/event/push/batch", h.PushBatchEventHandler)
	})
	srv := &http.Server{
		Addr:    cfg.HTTPServerAddress,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServeTLS("localhost.crt", "localhost.key"); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed initialize server")
		}
	}()
	return srv, connpool
}

// BuildApp - основная функция работы приложения.
func BuildApp() {
	log.Info().Msg("idm started")
	log.Info().Msg("Build version: " + buildVersion)
	log.Info().Msg("Build date: " + buildDate)
	log.Info().Msg("Build commit: " + buildCommit)
	cfg, err := configs.LoadAppConfig("./config/fanin.yaml")
	if err != nil {
		log.Fatal().Err(err).Msg("fail read env")
	}
	log.Info().Dict("cfg", zerolog.Dict().
		Str("HTTPServerAddress", cfg.HTTPServerAddress).
		Str("IdmAddress", cfg.IdmAddress).
		Str("CmdbAddress", cfg.CmdbAddress).
		Str("MetricsAddress", cfg.MetricsAddress).
		Str("EventsAddress", cfg.EventsAddress),
	).Msg("Server config")
	ctx := context.Background()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	srv, pool := HTTPServer(ctx, cfg)
	<-sigChan
	ctxcancel, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer func() {
		pool.Close()
		cancel()
	}()
	if err := srv.Shutdown(ctxcancel); err != nil {
		log.Error().Err(err).Msg("failed shutdown server")
	}
}
