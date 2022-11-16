package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/borisbbtest/GoMon/internal/mgrevent/configs"
	"github.com/borisbbtest/GoMon/internal/mgrevent/utils"
	"github.com/rs/zerolog/log"
)

type ServiceEvents struct {
}

var buildVersion = "N/A"
var buildDate = "N/A"
var buildCommit = "N/A"

func printIntro() {
	utils.Log.Debug().Msgf("Build version: ", buildVersion)
	utils.Log.Debug().Msgf("Build date: ", buildDate)
	utils.Log.Debug().Msgf("Build commit: ", buildCommit)
}

func Init(cfg *configs.MainConfig) (res *ServiceEvents, err error) {
	res = &ServiceEvents{}
	return res, nil
}

func (hook *ServiceEvents) Start() (err error) {
	utils.Log.Debug().Msg("Start Application events managers")

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
				utils.Log.Debug().Msg("bz -SIGINT")
				close(idleConnsClosed)
			case syscall.SIGTERM:
				if err := server.Shutdown(context.Background()); err != nil {
					// ошибки закрытия Listener
					log.Printf("HTTP server Shutdown SIGTERM: %v", err)
				}
				utils.Log.Debug().Msg("bz - SIGTERM")
				close(idleConnsClosed)
			case syscall.SIGQUIT:
				if err := server.Shutdown(context.Background()); err != nil {
					// ошибки закрытия Listener
					log.Printf("HTTP server Shutdown SIGQUIT: %v", err)
				}
				utils.Log.Debug().Msg("bz - SIGQUIT")
				close(idleConnsClosed)
			default:
				fmt.Println("Unknown signal.")
			}
		}
	}()

	<-idleConnsClosed

	return nil
}
