package models

import (
	"github.com/borisbbtest/GoMon/internal/idm/configs"
	"github.com/borisbbtest/GoMon/internal/idm/database"
	"github.com/borisbbtest/GoMon/internal/idm/service"
	"github.com/rs/zerolog"
)

var log = zerolog.New(service.LogConfig()).With().Timestamp().Caller().Logger()

type AppWrapper struct {
	Cfg  *configs.AppConfig
	Repo database.Storager
}
