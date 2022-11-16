package handler_grpc

import (
	"github.com/borisbbtest/GoMon/internal/mgrevent/configs"
	"github.com/borisbbtest/GoMon/internal/mgrevent/storage"
	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
)

type WrapperHandlerRPC struct {
	ServerConf *configs.MainConfig
	Storage    storage.Storage
	mgrevent.UnimplementedEventsServer
}
