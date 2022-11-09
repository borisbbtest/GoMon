package models

import (
	"github.com/rs/zerolog"

	"github.com/borisbbtest/GoMon/internal/idm/configs"
	"github.com/borisbbtest/GoMon/internal/idm/database"
	"github.com/borisbbtest/GoMon/internal/idm/service"
)

var log = zerolog.New(service.LogConfig()).With().Timestamp().Caller().Logger()

type ConfigWrapper struct {
	Cfg  *configs.AppConfig
	Repo database.Storager
}
