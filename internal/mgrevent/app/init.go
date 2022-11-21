// Package пакет по созданию приложения
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/borisbbtest/GoMon/internal/mgrevent/configs"
	servergrpc "github.com/borisbbtest/GoMon/internal/mgrevent/server"
	"github.com/borisbbtest/GoMon/internal/mgrevent/storage"
	storagepg "github.com/borisbbtest/GoMon/internal/mgrevent/storage/pg"
	"github.com/borisbbtest/GoMon/internal/mgrevent/utils"
)

// ServiceEvents класс приложения
type ServiceEvents struct {
	ServerConf *configs.MainConfig
	Storage    storage.Storage
}

var buildVersion = "N/A"
var buildDate = "N/A"
var buildCommit = "N/A"

func printIntro() {
	utils.Log.Debug().Msgf("Build version: ", buildVersion)
	utils.Log.Debug().Msgf("Build date: ", buildDate)
	utils.Log.Debug().Msgf("Build commit: ", buildCommit)
}

// Init - иницилизация  приложения
func Init(cfg *configs.MainConfig) (res *ServiceEvents, err error) {
	res = &ServiceEvents{}
	res.Storage, err = storagepg.NewPostgreSQLStorage(cfg)
	if err != nil {
		utils.Log.Debug().Err(err)
	}
	res.ServerConf = cfg
	return
}

// Start - запуск приложения
func (hook *ServiceEvents) Start() (err error) {
	utils.Log.Debug().Msg("Start Application events managers")

	go servergrpc.NewRPC(hook.ServerConf, hook.Storage).Start()
	// -------
	idleConnsClosed := make(chan struct{})
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	go func() {
		for {
			s := <-sigint
			switch s {
			case syscall.SIGINT:
				// if err := server.Shutdown(context.Background()); err != nil {
				// 	// ошибки закрытия Listener
				// 	log.Printf("HTTP server Shutdown SIGINT:  %v", err)
				// }
				utils.Log.Debug().Msg("bz -SIGINT")
				close(idleConnsClosed)
			case syscall.SIGTERM:
				// if err := server.Shutdown(context.Background()); err != nil {
				// 	// ошибки закрытия Listener
				// 	log.Printf("HTTP server Shutdown SIGTERM: %v", err)
				// }
				utils.Log.Debug().Msg("bz - SIGTERM")
				close(idleConnsClosed)
			case syscall.SIGQUIT:
				// if err := server.Shutdown(context.Background()); err != nil {
				// 	// ошибки закрытия Listener
				// 	log.Printf("HTTP server Shutdown SIGQUIT: %v", err)
				// }
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
