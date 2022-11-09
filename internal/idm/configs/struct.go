package configs

import (
	"time"

	"github.com/rs/zerolog"

	"github.com/borisbbtest/GoMon/internal/idm/service"
)

var log = zerolog.New(service.LogConfig()).With().Timestamp().Caller().Logger()

type AppConfig struct {
	DBDSN              string        `yaml:"DBDSN" env:"DATABASE_DSN"`              // URL для подключения к Postgres
	ServerAddressGRPC  string        `yaml:"ServerAddressGRPC" env:"ADDRESS_GRPC"`  // Адрес, по которому будут доступны endpoints
	SessionTimeExpired time.Duration `yaml:"SessionTimeExpired" env:"SESSION_TIME"` // Время жизни сессии в формате GO ("2m5s")
	ReInit             bool          `yaml:"ReInit" env:"REINIT"`                   // Требуется ли пересоздать таблицы в БД
}

func LoadAppConfig() *AppConfig {
	cfg := &AppConfig{
		DBDSN:              "postgres://pi:toor@192.168.1.69:5432/yandex",
		ServerAddressGRPC:  ":3200",
		SessionTimeExpired: 10 * time.Minute,
		ReInit:             true,
	}
	//yaml config
	cfg.yamlRead("./config/idm.yaml")
	//flags config
	cfg.flagsRead()
	//env config
	cfg.envRead()
	return cfg
}
