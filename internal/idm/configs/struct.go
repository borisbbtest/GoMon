// Package configs отвечает за получение конфига сервиса при старте.
package configs

import (
	"time"

	"github.com/rs/zerolog"

	"github.com/borisbbtest/GoMon/internal/idm/service"
)

var log = zerolog.New(service.LogConfig()).With().Timestamp().Caller().Logger()

// AppConfig - структура, описывающая параметры работы модуля idm
type AppConfig struct {
	DBDSN              string        `yaml:"DBDSN" env:"DATABASE_DSN"`              // URL для подключения к Postgres
	ServerAddressGRPC  string        `yaml:"ServerAddressGRPC" env:"ADDRESS_GRPC"`  // Адрес, по которому будут доступны endpoints
	SessionTimeExpired time.Duration `yaml:"SessionTimeExpired" env:"SESSION_TIME"` // Время жизни сессии в формате GO ("2m5s")
	ReInit             bool          `yaml:"ReInit" env:"REINIT"`                   // Требуется ли пересоздать таблицы в БД
}

// LoadAppConfig - создает AppConfig и заполняет его в следующем порядке:
//
// Значение по умолчанию -> yaml-файл -> переменные окружения -> флаги запуска.
//
// То, что находится правее в списке - будет в приоритете над тем, что левее.
func LoadAppConfig(file string) (*AppConfig, error) {
	cfg := &AppConfig{
		DBDSN:              "postgres://pi:toor@192.168.1.69:5432/yandex",
		ServerAddressGRPC:  ":3200",
		SessionTimeExpired: 10 * time.Minute,
		ReInit:             true,
	}
	//yaml config
	err := cfg.YamlRead(file)
	if err != nil {
		log.Error().Err(err).Msg("fail read yaml")
		return nil, err
	}
	//flags config
	cfg.FlagsRead()
	//env config
	err = cfg.EnvRead()
	if err != nil {
		log.Error().Err(err).Msg("fail read env")
		return nil, err
	}
	return cfg, nil
}
