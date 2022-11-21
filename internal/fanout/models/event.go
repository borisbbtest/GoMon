package models

import "time"

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
