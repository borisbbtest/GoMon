package grpc

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (hook *WrapperHandlerRPC) PushBatch(ctx context.Context, ev *mgrevent.PushBatchRequest) (res *mgrevent.PushBatchResponse, err error) {
	res = &mgrevent.PushBatchResponse{
		Code: "0",
	}
	e := ev.GetEv()
	err, err2 := hook.Storage.SaveEvents(ctx, e)
	if err != nil {
		return res, status.Error(codes.FailedPrecondition, "Opps OMG")
	}
	if err2 != nil {
		res.Code = err2.Error()
		return res, status.Error(codes.InvalidArgument, err2.Error())
	}

	return &mgrevent.PushBatchResponse{Code: "PushBatch"}, status.Error(codes.OK, "It is good response")
}
