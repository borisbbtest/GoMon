// Package models описывает бизнес-логику работы приложения с объектами системы
package models

import (
	"time"

	"github.com/borisbbtest/GoMon/internal/fanin/configs"
	"github.com/borisbbtest/GoMon/internal/fanin/service"
	"github.com/borisbbtest/GoMon/internal/models/cmdb"
	"github.com/borisbbtest/GoMon/internal/models/events"
	"github.com/borisbbtest/GoMon/internal/models/idm"
	"github.com/borisbbtest/GoMon/internal/models/metrics"
	"github.com/rs/zerolog"
)

var log = zerolog.New(service.LogConfig()).With().Timestamp().Caller().Logger()

// ClientPool - пул подключений для взаимодействия с компонентами системы при работе с бизнес-объектами
type ClientPool struct {
	Idm     idm.IdmClient         // клиент для модуля idm (управление пользователями)
	Cmdb    cmdb.CmdbClient       // клиент для модуля cmdb (управление КЕ)
	Metrics metrics.MetricsClient // клиент для модуля metrics (хранение метрик)
	Events  events.EventsClient   // клиент для модуля events (хранение событий)
}

// ConfigWrapper - структура конфигурации приложения
type ConfigWrapper struct {
	Cfg   *configs.AppConfig // конфигурационные параметры приложения
	Conns *ClientPool        // пулл соединений gRPC компонентов системы
}

// User - внутренний тип пользователя для данного модуля, используется при Unmarshall из входных данных HTTP
type User struct {
	Login     string    `json:"login"`               // имя пользователя, ключевой атрибут
	Firstname string    `json:"firstname,omitempty"` // имя
	Lastname  string    `json:"lastname,omitempty"`  // фамилия
	Password  string    `json:"password"`            // пароль
	CreatedAt time.Time `json:"createdat,omitempty"` // дата создания пользователя (заполняется системой)
	Source    string    `json:"source,omitempty"`    // источник пользователя
	Id        string    `json:"id,omitempty"`        // айди пользователя (заполняется системой)
}

// Session - внутренний тип сессии для данного модуля, используется при Unmarshall из входных данных HTTP
type Session struct {
	Id       string    `json:"id"`               // айди сессии (заполняется системой), ключевой атрибут
	Config   string    `json:"config,omitempty"` // параметр для хранения конфигураций конкретной сессии
	Login    string    `json:"login"`            // имя пользователя, ключевой атрибут
	Duration time.Time `json:"duration"`         // дата, когда сессия перестанет быть валидной (заполняется системой)
	Created  time.Time `json:"created"`          // дата создания сессии (заполняется системой)
	Code     string    `json:"code,omitempty"`   // код??
}

// Ci - внутренний тип КЕ для данного модуля, используется при Unmarshall из входных данных HTTP
type Ci struct {
	Name        string    `json:"name"`                  // название КЕ, ключевой атрибут
	Description string    `json:"description,omitempty"` // произвольное описание КЕ
	Update      time.Time `json:"update,omitempty"`      // дата обновления КЕ (заполняется системой)
	Created     time.Time `json:"created,omitempty"`     // дата создания КЕ (заполняется системой)
	CreatedBy   string    `json:"createdby,omitempty"`   // кем создана КЕ
	Type        string    `json:"type"`                  // тип КЕ
}

// Metric - внутренний тип метрики для данного модуля, используется при Unmarshall из входных данных HTTP
type Metric struct {
	Name              string    `json:"name"`      // имя метрики
	Value             []byte    `json:"value"`     // значение метрики
	Localtime         time.Time `json:"localtime"` // дата этой метрики
	SourceFromSystems string    `json:"source"`    // система источник метрики для случая, когда разные источники могут прислать одну и ту же метрику
	RelationCi        string    `json:"ci"`        // КЕ, к которой относится эта метрика
}

// Event - внутренний тип события для данного модуля, используется при Unmarshall из входных данных HTTP
type Event struct {
	Title       string    `json:"title"`                 // краткий заголовок события
	Description string    `json:"description,omitempty"` // подробное описание события
	Source      string    `json:"source,omitempty"`      // источник события
	Status      string    `json:"status"`                // статус события
	Created     time.Time `json:"created,omitempty"`     // дата создания события
	Update      time.Time `json:"update,omitempty"`      // дата обновления события
	Key         string    `json:"key,omitempty"`         // ключ события для корреляции
	KeyClose    string    `json:"keyclose,omitempty"`    // паттерн ключа, который закроет данное событие
	Assigned    string    `json:"assigned,omitempty"`    // на кого назначено событие
	AutoRunner  string    `json:"autorunner,omitempty"`  // описание автоматического действия по событию
	Severity    string    `json:"severity"`              // критичность события
	RelarionCi  string    `json:"relationci"`            // к какое КЕ относится это событие
	CreatedBy   string    `json:"createdby,omitempty"`   // кем создано событие
	Count       int32     `json:"count,omitempty"`       // количество пришедших одинаковых событий (количество дублей)
	Id          string    `json:"id,omitempty"`          // айди события
}

// FanInContextKey - тип, использующийся при создании context withValue для избежания возможных коллизий
type FanInContextKey string
