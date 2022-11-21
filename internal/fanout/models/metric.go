package models

import "time"

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
