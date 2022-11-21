package grpc

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/metrics/utils"
	"github.com/borisbbtest/GoMon/internal/models/metrics"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

// GetBatch получает список метрик
func (hook *WrapperHandlerRPC) GetBatch(ctx context.Context, item *metrics.GetBatchRequest) (res *metrics.GetBatchResponse, err error) {

	utils.Log.Info().Msg("GetBatch start")
	res = &metrics.GetBatchResponse{
		Code: "Ok",
	}

	err, tmpeve := hook.Storage.GetMetricsDuration(ctx, item.Start.AsTime(), item.End.AsTime())
	if err != nil {
		utils.Log.Error().Msgf(err.Error())
		return res, status.Error(codes.FailedPrecondition, "Opps OMG")
	}
	buff := tmpeve.ConvertTogRpcEvent()

	for _, v := range *buff {
		res.Item = append(res.Item, v)
	}

	return res, nil

}
