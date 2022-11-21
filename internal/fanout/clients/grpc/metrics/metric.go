package integrationmetric

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/borisbbtest/GoMon/internal/fanin/service"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
	"github.com/borisbbtest/GoMon/internal/models/metrics"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Metric - внутренний тип метрики для данного модуля, используется при Unmarshall из входных данных HTTP
type Metric struct {
	Name              string    `json:"name"`                  // имя метрики
	Value             []byte    `json:"value"`                 // значение метрики
	Localtime         time.Time `json:"localtime"`             // дата этой метрики по загрузке
	SourceTime        time.Time `json:"source_time,omitempty"` // дата этой метрики от источника
	SourceFromSystems string    `json:"source_from_systems"`   // система источник метрики для случая, когда разные источники могут прислать одну и ту же метрику
	RelationCi        string    `json:"relation_ci"`           // КЕ, к которой относится эта метрика
	Uuid              string    `json:"uuid,omitempty"`        // id метрики
	Tp                string    `json:"tp"`                    // тип метрики
}

// textToTypes - парсинг string в protobuf Types
func textToTypes(types string) (metrics.Types, error) {
	if v, ok := metrics.Types_value[types]; ok {
		pbstatus := metrics.Types(v)
		return pbstatus, nil
	}
	inttypes, err := strconv.ParseInt(types, 10, 32)
	if err != nil {
		return metrics.Types(0), err
	}
	if _, ok := metrics.Types_name[int32(inttypes)]; ok {
		pbstatus := metrics.Types(int32(inttypes))
		return pbstatus, nil
	}
	return metrics.Types(0), service.ErrMetricWrongType
}

// ToPB - конвертация типа Metric в protobuf  Metric
func (hook *Metric) ToPB() (*metrics.Metric, error) {
	tp, err := textToTypes(hook.Tp)
	if err != nil {
		return nil, err
	}
	return &metrics.Metric{
		Name:              hook.Name,
		Value:             hook.Value,
		Localtime:         timestamppb.New(hook.Localtime),
		SourceTime:        timestamppb.New(hook.SourceTime),
		SourceFromSystems: hook.SourceFromSystems,
		RelationCi:        hook.RelationCi,
		Uuid:              hook.Uuid,
		Tp:                tp,
	}, nil
}

// UnmarshalJSON - функция переопределяющия правила анмаршалера для timestamp в Metric
func (hook *Metric) UnmarshalJSON(data []byte) error {
	type MetricAlias Metric
	AliasValue := &struct {
		*MetricAlias
		Localtime  string `json:"localtime"`
		SourceTime string `json:"source_time"`
	}{
		MetricAlias: (*MetricAlias)(hook),
	}
	if err := json.Unmarshal(data, AliasValue); err != nil {
		utils.Log.Error().Err(err).Msg("failed unmarshall json")
		return err
	}
	if AliasValue.Localtime != "" {
		res, err := time.Parse(time.RFC3339, AliasValue.Localtime)
		if err != nil {
			utils.Log.Error().Err(err).Msg("failed unmarshall Localtime")
			return err
		}
		hook.Localtime = res
	}
	if AliasValue.SourceTime != "" {
		res, err := time.Parse(time.RFC3339, AliasValue.SourceTime)
		if err != nil {
			utils.Log.Error().Err(err).Msg("failed unmarshall SourceTime")
			return err
		}
		hook.SourceTime = res
	}
	return nil
}
