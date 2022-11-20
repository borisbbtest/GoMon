package grpc

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/models/metrics"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (hook *WrapperHandlerRPC) PushBatch(ctx context.Context, ev *metrics.PushBatchRequest) (res *metrics.PushBatchResponse, err error) {
	res = &metrics.PushBatchResponse{
		Code: "0",
	}
	e := ev.Item
	err, err2 := hook.Storage.SaveEvents(ctx, e)
	if err != nil {
		return res, status.Error(codes.FailedPrecondition, "Opps OMG")
	}
	if err2 != nil && err2.Error() != "no result" {
		res.Code = err2.Error()
		return res, status.Error(codes.InvalidArgument, err2.Error())
	}

	return &metrics.PushBatchResponse{Code: "PushBatch"}, status.Error(codes.OK, "It is good response")
}
