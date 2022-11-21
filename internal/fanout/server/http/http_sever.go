package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/borisbbtest/GoMon/internal/fanout/configs"
	handlers_http "github.com/borisbbtest/GoMon/internal/fanout/handlers/http"
	midd "github.com/borisbbtest/GoMon/internal/fanout/middleware"
	"github.com/borisbbtest/GoMon/internal/fanout/models"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

// serviceHTTPFanOut Класс сервера htttp
type serviceHTTPFanOut struct {
	wrapp  handlers_http.WrapperHandler
	middle midd.WrapperMiddleware
}

// NewHTTP конструктор класс serviceHTTPFanOut
func NewHTTP(cfg *config.MainConfig) *serviceHTTPFanOut {

	init := &models.ServicePoolWrapper{}
	_ = init

	return &serviceHTTPFanOut{
		wrapp: handlers_http.WrapperHandler{
			ServerConf: cfg,
		},
		middle: midd.WrapperMiddleware{},
	}
}

var buildVersion = "N/A"
var buildDate = "N/A"
var buildCommit = "N/A"

func printIntro() {
	utils.Log.Debug().Msgf("Build version: ", buildVersion)
	utils.Log.Debug().Msgf("Build date: ", buildDate)
	utils.Log.Debug().Msgf("Build commit: ", buildCommit)
}

// Start Запускает сервер http
func (hook *serviceHTTPFanOut) Start() (err error) {

	printIntro()
	// Launch the listening thread
	log.Println("Initializing HTTP server")
	r := chi.NewRouter()

	//	defer hook.wrapp.Storage.Close()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(midd.GzipHandle)
	r.Use(middleware.Recoverer)
	r.HandleFunc("/pprof/*", pprof.Index)
	r.HandleFunc("/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/pprof/profile", pprof.Profile)
	r.HandleFunc("/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/pprof/trace", pprof.Trace)
	r.Handle("/pprof/goroutine", pprof.Handler("goroutine"))
	r.Handle("/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Handle("/pprof/mutex", pprof.Handler("mutex"))
	r.Handle("/pprof/heap", pprof.Handler("heap"))
	r.Handle("/pprof/block", pprof.Handler("block"))
	r.Handle("/pprof/allocs", pprof.Handler("allocs"))

	r.Get("/", hook.wrapp.PingHandler)
	r.Post("/api/register", hook.wrapp.RegisterHandler)
	r.Post("/api/authorize", hook.wrapp.AuthorizeHandler)

	serviceLogic := r.Group(nil)
	serviceLogic.Use(hook.middle.MiddleSetSessionCookie)
	// CMD
	serviceLogic.Get("/api/get_ci/{name}", hook.wrapp.GetGetCi)
	serviceLogic.Post("/api/get_ci", hook.wrapp.PostGetCis)
	//Events
	serviceLogic.Get("/api/get_event/{id}", hook.wrapp.GetGetEvent)
	serviceLogic.Post("/api/get_events", hook.wrapp.PostGetEvens)
	//Metric
	serviceLogic.Post("/api/get_metrics", hook.wrapp.PostGetMetrics)

	server := &http.Server{
		Addr:         hook.wrapp.ServerConf.AccrualSystemAddress,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 40 * time.Second,
	}

	// через этот канал сообщим основному потоку, что соединения закрыты
	idleConnsClosed := make(chan struct{})
	// канал для перенаправления прерываний
	// поскольку нужно отловить всего одно прерывание,
	// ёмкости 1 для канала будет достаточно
	sigint := make(chan os.Signal, 1)
	// регистрируем перенаправление прерываний
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	// запускаем горутину обработки пойманных прерываний
	go func() {
		// читаем из канала прерываний
		// поскольку нужно прочитать только одно прерывание,
		// можно обойтись без цикла
		for {
			s := <-sigint
			switch s {
			case syscall.SIGINT:
				if err := server.Shutdown(context.Background()); err != nil {
					// ошибки закрытия Listener
					log.Printf("HTTP server Shutdown SIGINT:  %v", err)
				}
				utils.Log.Debug().Msgf("bz -SIGINT")
				close(idleConnsClosed)
			case syscall.SIGTERM:
				if err := server.Shutdown(context.Background()); err != nil {
					// ошибки закрытия Listener
					utils.Log.Debug().Msgf("HTTP server Shutdown SIGTERM: %v", err)
				}
				utils.Log.Debug().Msgf("bz - SIGTERM")
				close(idleConnsClosed)
			case syscall.SIGQUIT:
				if err := server.Shutdown(context.Background()); err != nil {
					// ошибки закрытия Listener
					utils.Log.Debug().Msgf("HTTP server Shutdown SIGQUIT: %v", err)
				}
				utils.Log.Debug().Msgf("bz - SIGQUIT")
				close(idleConnsClosed)
			default:
				fmt.Println("Unknown signal.")
			}
		}
	}()

	defer server.Close()
	if hook.wrapp.ServerConf.EnableHTTPS {

		cert, key, err := utils.CertGeg()
		if err != nil {
			return fmt.Errorf("BZ Certificate and key wasn't generation: %s", err)
		}

		utils.WriteCertFile("cert.pem", cert)
		utils.WriteCertFile("key.pem", key)
		err = server.ListenAndServeTLS("cert.pem", "key.pem")

		if err != http.ErrServerClosed {
			return fmt.Errorf("BZ can't start the listening thread: %s", err)
		}

	} else {
		err = server.ListenAndServe()
		if err != http.ErrServerClosed {

			return fmt.Errorf("BZ can't start the listening thread: %s", err)
		}

	}
	// ждём завершения процедуры graceful shutdown
	<-idleConnsClosed
	// получили оповещение о завершении
	// здесь можно освобождать ресурсы перед выходом,
	// например закрыть соединение с базой данных,
	// закрыть открытые файлы

	utils.Log.Debug().Msgf("Server Shutdown gracefully")

	utils.Log.Debug().Msgf("Exiting")
	return nil
}
