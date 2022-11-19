package grpc

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
)

func (hook *WrapperHandlerRPC) Push(context.Context, *mgrevent.PushRequest) (*mgrevent.PushResponse, error) {

	return &mgrevent.PushResponse{Code: "Push"}, nil
}
