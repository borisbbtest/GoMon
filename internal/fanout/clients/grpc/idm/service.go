// Package models описывает бизнес-логику работы приложения с объектами системы
package integrationidm

import (
	"github.com/borisbbtest/GoMon/internal/models/idm"
)

// ServiceWrapperIdm - класс по работе с сервисом IDM
type ServiceWrapperIdm struct {
	Idm idm.IdmClient //соединений gRPC IDM
}
