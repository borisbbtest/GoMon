package grpc

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
)

func (hook *WrapperHandlerRPC) GetBatch(context.Context, *mgrevent.GetBatchRequest) (*mgrevent.GetBatchResponse, error) {

	return &mgrevent.GetBatchResponse{Code: "GetBatch"}, nil
}
