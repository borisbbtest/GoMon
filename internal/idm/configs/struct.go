package configs

import (
	"github.com/rs/zerolog"

	"github.com/borisbbtest/GoMon/internal/idm/service"
)

var log = zerolog.New(service.LogConfig()).With().Timestamp().Caller().Logger()

type AppConfig struct {
	DBDSN             string `env:"DATABASE_DSN"` // URL для подключения к Postgres
	ServerAddressGRPC string `env:"ADDRESS_GRPC"` // Адрес, по которому будут доступны endpoints
	ReInit            bool   `env:"REINIT"`       // Требуется ли пересоздать таблицы в БД
}

func LoadAppConfig() *AppConfig {
	cfg := &AppConfig{
		DBDSN:             "postgres://pi:toor@192.168.1.69:5432/yandex",
		ServerAddressGRPC: ":3200",
		ReInit:            true,
	}
	//flags config
	cfg.flagsRead()
	//env config
	cfg.envRead()
	return cfg
}
