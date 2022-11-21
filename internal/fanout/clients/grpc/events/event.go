package integrationevents

import (
	"encoding/json"
	"time"

	"github.com/rs/zerolog/log"
)

// Event - внутренний тип события для данного модуля, используется при Unmarshall из входных данных HTTP
type Event struct {
	Title       string    `json:"title"`                 // краткий заголовок события
	Description string    `json:"description,omitempty"` // подробное описание события
	Source      string    `json:"source,omitempty"`      // источник события
	Status      int32     `json:"status"`                // статус события
	Created     time.Time `json:"created,omitempty"`     // дата создания события
	Update      time.Time `json:"update,omitempty"`      // дата обновления события
	Key         string    `json:"key,omitempty"`         // ключ события для корреляции
	KeyClose    string    `json:"key_close,omitempty"`   // паттерн ключа, который закроет данное событие
	Assigned    []string  `json:"assigned,omitempty"`    // на кого назначено событие
	AutoRunner  string    `json:"auto_runner,omitempty"` // описание автоматического действия по событию
	Severity    int32     `json:"severity"`              // критичность события
	RelarionCi  []string  `json:"relarion_ci"`           // к какое КЕ относится это событие
	CreatedBy   string    `json:"created_by,omitempty"`  // кем создано событие
	Count       int32     `json:"count,omitempty"`       // количество пришедших одинаковых событий (количество дублей)
	Uuid        string    `json:"uuid,omitempty"`        // айди события
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type EventAlias Event
	AliasValue := &struct {
		*EventAlias
		Created string `json:"created"`
		Update  string `json:"update"`
	}{
		EventAlias: (*EventAlias)(e),
	}
	if err := json.Unmarshal(data, AliasValue); err != nil {
		log.Error().Err(err).Msg("failed unmarshall json")
		return err
	}
	if AliasValue.Created != "" {
		res, err := time.Parse(time.RFC3339, AliasValue.Created)
		if err != nil {
			log.Error().Err(err).Msg("failed unmarshall Created")
			return err
		}
		e.Created = res
	}
	if AliasValue.Update != "" {
		res, err := time.Parse(time.RFC3339, AliasValue.Update)
		if err != nil {
			log.Error().Err(err).Msg("failed unmarshall Update")
			return err
		}
		e.Update = res
	}
	return nil
}
