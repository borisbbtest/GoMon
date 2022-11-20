package app

import (
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
	"github.com/borisbbtest/go_home_work/internal/config"
	"github.com/borisbbtest/go_home_work/internal/storage"
	"github.com/rs/zerolog/log"
)

type ServiceShortURL struct {
	ServerConf *config.ServiceFanOutConfig
	Storage    storage.Storage
}

func Init(cfg *config.ServiceShortURLConfig) (res *ServiceShortURL, err error) {

	res = &ServiceShortURL{}
	res.Storage, err = storage.NewPostgreSQLStorage(cfg.DataBaseDSN)
	if err != nil {
		res.Storage, err = storage.NewFileStorage(cfg.FileStorePath)
		if err != nil {
			utils.Log.Error().Err(err)
		}
	}
	res.ServerConf = cfg
	return
}

func (hook *ServiceShortURL) Start() (err error) {

	// log.Info("Start RPC")
	// go NewRPC(hook.ServerConf, hook.Storage).Start()

	utils.Log.Info().Msgf("Start HTTP")
	err = NewHTTP(hook.ServerConf, hook.Storage).Start()
	if err != nil {
		log.Fatal(err)
		return
	}
	return

}
