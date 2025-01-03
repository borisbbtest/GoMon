package grpc

import (
	"github.com/borisbbtest/GoMon/internal/mgrevent/configs"
	"github.com/borisbbtest/GoMon/internal/mgrevent/storage"
	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
)

// WrapperHandlerRPC - класс обработчик grpc запросов
type WrapperHandlerRPC struct {
	ServerConf *configs.MainConfig
	Storage    storage.Storage
	mgrevent.UnimplementedEventsServer
}
