package grpc

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/mgrevent/utils"
	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (hook *WrapperHandlerRPC) GetBatch(ctx context.Context, eve *mgrevent.GetBatchRequest) (res *mgrevent.GetBatchResponse, err error) {

	utils.Log.Info().Msg("GetBatch start")
	res = &mgrevent.GetBatchResponse{
		Code: "0",
	}

	err, tmpeve := hook.Storage.GetEventDuration(ctx, eve.Start.AsTime(), eve.End.AsTime())
	if err != nil {
		utils.Log.Error().Msgf(err.Error())
		return res, status.Error(codes.FailedPrecondition, "Opps OMG")
	}
	buff := tmpeve.ConvertTogRpcEvent()

	for _, v := range *buff {
		res.Ev = append(res.Ev, v)
	}

	return res, nil

}
