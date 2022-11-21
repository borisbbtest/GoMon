// Package configs отвечает за получение конфига сервиса при старте.
package configs

import (
	"github.com/rs/zerolog"

	"github.com/borisbbtest/GoMon/internal/fanin/service"
)

var log = zerolog.New(service.LogConfig()).With().Timestamp().Caller().Logger()

// AppConfig - структура, описывающая параметры работы модуля cmdb
type AppConfig struct {
	HTTPServerAddress string `yaml:"HTTPServerAddress" env:"HTTP_ADDRESS"` // Адрес сервера, на котором будут ожидатся подключения
	IdmAddress        string `yaml:"IdmAddress" env:"IDM_ADDRESS"`         // Адрес модуля idm
	CmdbAddress       string `yaml:"CmdbAddress" env:"CMDB_ADDRESS"`       // Адрес модуля cmdb
	MetricsAddress    string `yaml:"MetricsAddress" env:"METRICS_ADDRESS"` // Адрес модуля metrics
	EventsAddress     string `yaml:"EventsAddress" env:"EVENTS_ADDRESS"`   // Адрес модуля mgrevent
}

// LoadAppConfig - создает AppConfig и заполняет его в следующем порядке:
//
// Значение по умолчанию -> yaml-файл -> флаги запуска -> переменные окружения
//
// То, что находится правее в списке - будет в приоритете над тем, что левее.
func LoadAppConfig(file string) (*AppConfig, error) {
	cfg := &AppConfig{
		HTTPServerAddress: ":443",
		IdmAddress:        ":8081",
		CmdbAddress:       ":8082",
		MetricsAddress:    ":8083",
		EventsAddress:     ":8084",
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
