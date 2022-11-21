package models

import (
	"encoding/json"
	"time"

	integrationevents "github.com/borisbbtest/GoMon/internal/fanout/clients/grpc/events"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
)

type RequestGetEvent struct {
	StartTime time.Time
	EndTime   time.Time
	Uuid      string
}

type ResponseGetEventDuration struct {
	Root *[]*integrationevents.Event
}

type ResponseGetEvent struct {
	Root *integrationevents.Event
}

// ParseRequestDuration - функция переопределяющия правила анмаршалера для timestamp в Metric
func (hook *RequestGetEvent) ParseRequestDuration(data []byte) error {
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
