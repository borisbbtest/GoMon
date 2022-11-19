package grpc

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
)

func (hook *WrapperHandlerRPC) PushBatch(context.Context, *mgrevent.PushBatchRequest) (*mgrevent.PushBatchResponse, error) {

	return &mgrevent.PushBatchResponse{Code: "PushBatch"}, nil
}
