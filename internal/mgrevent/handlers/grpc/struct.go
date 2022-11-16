package grpc

import "github.com/prometheus/prometheus/storage"

type WrapperHandlerRPC struct {
	ServerConf *config.ServiceShortURLConfig
	Storage    storage.Storage
	UserID     string
	shortrpc.UnimplementedShortURLServer
}
