package integrationcmdb

import "time"

// Ci - внутренний тип КЕ для данного модуля, используется при Unmarshall из входных данных HTTP
type Ci struct {
	Name        string    `json:"name"`                  // название КЕ, ключевой атрибут
	Description string    `json:"description,omitempty"` // произвольное описание КЕ
	Update      time.Time `json:"update,omitempty"`      // дата обновления КЕ (заполняется системой)
	Created     time.Time `json:"created,omitempty"`     // дата создания КЕ (заполняется системой)
	CreatedBy   string    `json:"created_by,omitempty"`  // кем создана КЕ
	Type        string    `json:"type"`                  // тип КЕ
}
