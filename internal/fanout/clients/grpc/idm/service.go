// Package models описывает бизнес-логику работы приложения с объектами системы
package integrationidm

import (
	"github.com/borisbbtest/GoMon/internal/models/idm"
)

// ConfigWrapper - структура конфигурации приложения
type ServiceWrapperIdm struct {
	Idm idm.IdmClient //соединений gRPC IDM
}
