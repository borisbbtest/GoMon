package grpc

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
)

func (hook *WrapperHandlerRPC) Get(context.Context, *mgrevent.GetRequest) (*mgrevent.GetResponse, error) {

	return &mgrevent.GetResponse{Code: "Get"}, nil
}
