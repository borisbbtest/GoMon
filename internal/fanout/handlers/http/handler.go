package handlers_http

import (
	config "github.com/borisbbtest/GoMon/internal/fanout/configs"
	"github.com/borisbbtest/GoMon/internal/fanout/models"
)

// WrapperHandler - класс хедлеров
type WrapperHandler struct {
	ServerConf  *config.MainConfig // Конфигурация приложения
	ServicePool *models.ClientPool // Сессия пользователя
}
