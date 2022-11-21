package grpc

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/mgrevent/utils"
	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (hook *WrapperHandlerRPC) Get(ctx context.Context, ev *mgrevent.GetRequest) (res *mgrevent.GetResponse, err error) {
	utils.Log.Info().Msg("Get start")
	res = &mgrevent.GetResponse{
		Code: "Ok",
	}
	e := &mgrevent.Event{Uuid: ev.GetId()}

	err, tmpeve := hook.Storage.GetEvent(ctx, e)
	if err != nil {
		utils.Log.Error().Msgf(err.Error())
		return nil, status.Error(codes.NotFound, " Get Opps OMG")
	}
	buff := tmpeve.ConvertTogRpcEvent()

	for _, v := range *buff {
		res.Ev = v
		return res, nil
	}

	return &mgrevent.GetResponse{Code: "Get"}, status.Error(codes.NotFound, "didn't find")
}
