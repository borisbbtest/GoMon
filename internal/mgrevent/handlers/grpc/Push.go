package grpc

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/metrics/utils"
	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (hook *WrapperHandlerRPC) Push(ctx context.Context, ev *mgrevent.PushRequest) (res *mgrevent.PushResponse, err error) {

	utils.Log.Debug().Msg("Psuh")
	res = &mgrevent.PushResponse{
		Code: "Push 0",
	}
	var x []*mgrevent.Event
	e := ev.GetEv()
	x = append(x, e)
	err, err2 := hook.Storage.SaveEvents(ctx, x)
	if err != nil {
		return res, status.Error(codes.FailedPrecondition, "Opps OMG")
	}
	if err2 != nil && err2.Error() != "no result" {
		res.Code = err2.Error()
		return res, status.Error(codes.InvalidArgument, err2.Error())
	}

	return &mgrevent.PushResponse{Code: "Push"}, status.Error(codes.OK, "It is good response")
}
