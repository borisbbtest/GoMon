package models

import (
	"encoding/json"
	"time"

	pb "github.com/borisbbtest/GoMon/internal/models/metrics"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToPB - конвертация типа Metric в protobuf  Metric
func (m *Metric) ToPB() *pb.Metric {
	return &pb.Metric{
		Name:              m.Name,
		Value:             m.Value,
		Localtime:         timestamppb.New(m.Localtime),
		SourceFromSystems: m.SourceFromSystems,
		RelationCi:        m.RelationCi,
	}
}

// UnmarshalJSON - функция переопределяющия правила анмаршалера для timestamp в Metric
func (m *Metric) UnmarshalJSON(data []byte) error {
	type MetricAlias Metric
	AliasValue := &struct {
		*MetricAlias
		Localtime string `json:"localtime"`
	}{
		MetricAlias: (*MetricAlias)(m),
	}
	if err := json.Unmarshal(data, AliasValue); err != nil {
		log.Error().Err(err).Msg("failed unmarshall json")
		return err
	}
	if AliasValue.Localtime != "" {
		res, err := time.Parse(time.RFC3339, AliasValue.Localtime)
		if err != nil {
			log.Error().Err(err).Msg("failed unmarshall Localtime")
			return err
		}
		m.Localtime = res
	}
	return nil
}
