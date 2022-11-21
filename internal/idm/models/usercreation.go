package models

import (
	"context"

	pb "github.com/borisbbtest/GoMon/internal/models/idm"
)

// CreateUser - функция создания пользователя в хранилище
func (w *ConfigWrapper) CreateUser(ctx context.Context, user *pb.User) error {
	err := w.Repo.CreateUser(ctx, w.Cfg, user)
	if err != nil {
		log.Error().Err(err).Msg("failed create user in db")
		return err
	}
	return nil
}
