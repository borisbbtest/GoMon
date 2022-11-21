package handlers_http

import (
	config "github.com/borisbbtest/GoMon/internal/fanout/configs"
	"github.com/borisbbtest/GoMon/internal/fanout/storage"
)

// WrapperHandler - класс хедлеров
type WrapperHandler struct {
	ServerConf *config.MainConfig   // Конфигурация приложения
	Storage    storage.Storage      // Ссылка на конект к бд
	Session    *storage.SessionHTTP // Сессия пользователя
}
