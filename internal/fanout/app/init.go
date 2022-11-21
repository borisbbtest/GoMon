package app

import (
	config "github.com/borisbbtest/GoMon/internal/fanout/configs"
	"github.com/borisbbtest/GoMon/internal/fanout/server/http"
	"github.com/borisbbtest/GoMon/internal/fanout/storage"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
)

// ServiceShortURL- Сервис приложения
type ServiceShortURL struct {
	ServerConf *config.MainConfig // Конфиг сервера и приложения
	Storage    storage.Storage    // Хранилка приложения
}

// Init- конструктор приложения
func Init(cfg *config.MainConfig) (res *ServiceShortURL, err error) {

	res = &ServiceShortURL{}
	res.Storage, err = storage.NewPostgreSQLStorage(cfg.DatabaseURI)
	if err != nil {
		utils.Log.Error().Err(err)
	}
	res.ServerConf = cfg
	return
}

// Start- запуск приложения
func (hook *ServiceShortURL) Start() (err error) {

	// log.Info("Start RPC")
	// go NewRPC(hook.ServerConf, hook.Storage).Start()

	utils.Log.Info().Msgf("Start HTTP")
	err = http.NewHTTP(hook.ServerConf, hook.Storage).Start()
	if err != nil {
		utils.Log.Error().Err(err)
		return
	}
	return

}
