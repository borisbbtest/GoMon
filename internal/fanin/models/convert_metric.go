package models

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/borisbbtest/GoMon/internal/fanin/service"
	pb "github.com/borisbbtest/GoMon/internal/models/metrics"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// textToTypes - парсинг string в protobuf Types
func textToTypes(types string) (pb.Types, error) {
	if v, ok := pb.Types_value[types]; ok {
		pbstatus := pb.Types(v)
		return pbstatus, nil
	}
	inttypes, err := strconv.ParseInt(types, 10, 32)
	if err != nil {
		return pb.Types(0), err
	}
	if _, ok := pb.Types_name[int32(inttypes)]; ok {
		pbstatus := pb.Types(int32(inttypes))
		return pbstatus, nil
	}
	return pb.Types(0), service.ErrMetricWrongType
}

// ToPB - конвертация типа Metric в protobuf  Metric
func (m *Metric) ToPB() (*pb.Metric, error) {
	tp, err := textToTypes(m.Tp)
	if err != nil {
		return nil, err
	}
	return &pb.Metric{
		Name:              m.Name,
		Value:             m.Value,
		Localtime:         timestamppb.New(m.Localtime),
		SourceTime:        timestamppb.New(m.SourceTime),
		SourceFromSystems: m.SourceFromSystems,
		RelationCi:        m.RelationCi,
		Uuid:              m.Uuid,
		Tp:                tp,
	}, nil
}

// UnmarshalJSON - функция переопределяющия правила анмаршалера для timestamp в Metric
func (m *Metric) UnmarshalJSON(data []byte) error {
	type MetricAlias Metric
	AliasValue := &struct {
		*MetricAlias
		Localtime  string `json:"localtime"`
		SourceTime string `json:"source_time"`
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
	if AliasValue.SourceTime != "" {
		res, err := time.Parse(time.RFC3339, AliasValue.SourceTime)
		if err != nil {
			log.Error().Err(err).Msg("failed unmarshall SourceTime")
			return err
		}
		m.SourceTime = res
	}
	return nil
}
