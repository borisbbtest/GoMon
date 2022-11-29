package grpc

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

// PushBatch отправка нескольких событий
func (hook *WrapperHandlerRPC) PushBatch(ctx context.Context, ev *mgrevent.PushBatchRequest) (res *mgrevent.PushBatchResponse, err error) {
	res = &mgrevent.PushBatchResponse{
		Code: "Ok",
	}
	e := ev.GetEv()
	err, err2 := hook.Storage.SaveEvents(ctx, e)
	if err != nil {
		return res, status.Error(codes.FailedPrecondition, "Opps OMG")
	}
	if err2 != nil && err2.Error() != "no result" {
		res.Code = err2.Error()
		return res, status.Error(codes.InvalidArgument, err2.Error())
	}

	return &mgrevent.PushBatchResponse{Code: "OK"}, nil
}
