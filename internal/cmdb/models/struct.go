// Package models описывает бизнес логику приложения по работе с КЕ
package models

import (
	"github.com/rs/zerolog"

	"github.com/borisbbtest/GoMon/internal/cmdb/configs"
	"github.com/borisbbtest/GoMon/internal/cmdb/database"
	"github.com/borisbbtest/GoMon/internal/cmdb/service"
)

var log = zerolog.New(service.LogConfig()).With().Timestamp().Caller().Logger()

// ConfigWrapper - структура конфигурации приложения
type ConfigWrapper struct {
	Cfg  *configs.AppConfig
	Repo database.Storager
}
