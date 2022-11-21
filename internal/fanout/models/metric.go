package models

import (
	"encoding/json"
	"time"

	integrationmetric "github.com/borisbbtest/GoMon/internal/fanout/clients/grpc/metrics"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
)

type RequestGetMetricDuration struct {
	StartTime time.Time
	EndTime   time.Time
}

type ResponseGetMetricDuration struct {
	Root *[]*integrationmetric.Metric
}

// UnmarshalJSON - функция переопределяющия правила анмаршалера для timestamp в Metric
func (hook *RequestGetMetricDuration) ParseRequest(data []byte) error {
	Req := &struct {
		StartTime string `json:"start"`
		EndTime   string `json:"end"`
	}{}
	if err := json.Unmarshal(data, Req); err != nil {
		utils.Log.Error().Err(err).Msg("failed unmarshall json")
		return err
	}
	if Req.StartTime != "" {
		res, err := time.Parse(time.RFC3339, Req.StartTime)
		if err != nil {
			utils.Log.Error().Err(err).Msg("failed unmarshall Localtime")
			return err
		}
		hook.StartTime = res
	}
	if Req.EndTime != "" {
		res, err := time.Parse(time.RFC3339, Req.EndTime)
		if err != nil {
			utils.Log.Error().Err(err).Msg("failed unmarshall SourceTime")
			return err
		}
		hook.EndTime = res
	}
	return nil
}
