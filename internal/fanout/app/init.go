package app

import (
	config "github.com/borisbbtest/GoMon/internal/fanout/configs"
	"github.com/borisbbtest/GoMon/internal/fanout/server/http"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
)

// ServiceShortURL- Сервис приложения
type ServiceFanIn struct {
	ServerConf *config.MainConfig // Конфиг сервера и приложения
}

// Init- конструктор приложения
func Init(cfg *config.MainConfig) (res *ServiceFanIn, err error) {
	res = &ServiceFanIn{}
	res.ServerConf = cfg
	return
}

// Start- запуск приложения
func (hook *ServiceFanIn) Start() (err error) {

	// log.Info("Start RPC")
	// go NewRPC(hook.ServerConf, hook.Storage).Start()

	utils.Log.Info().Msgf("Start HTTP")
	err = http.NewHTTP(hook.ServerConf).Start()
	if err != nil {
		utils.Log.Error().Err(err)
		return
	}
	return

}
