// Package models описывает бизнес-логику работы приложения с объектами системы
package integration_idm

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/models/idm"
)

// ConfigWrapper - структура конфигурации приложения
type ServiceWrapperIdm struct {
	Idm idm.IdmClient //соединений gRPC IDM
}
type ClientIdm interface {
	AuthorizeUser(ctx context.Context, user *User) (Session, error)
	CheckAuthorized(ctx context.Context, login string, id string) bool
	RegisterUser(ctx context.Context, user *User) (Session, error)
	Close()
}
