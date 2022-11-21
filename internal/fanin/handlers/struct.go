// Package handlers описывает endpoints для пользователей
package handlers

import (
	"github.com/borisbbtest/GoMon/internal/fanin/models"
	"github.com/borisbbtest/GoMon/internal/fanin/service"
	"github.com/rs/zerolog"
)

var log = zerolog.New(service.LogConfig()).With().Timestamp().Caller().Logger()

// HTTP - структура, методы которой будут являться хэндлерами сервера
type HTTP struct {
	App *models.ConfigWrapper // структура, хранящая пул подключений и конфиг приложения
}
