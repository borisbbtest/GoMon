package models

import (
	"context"

	pb "github.com/borisbbtest/GoMon/internal/idm/proto/idm"
)

func (w *ConfigWrapper) CreateUser(ctx context.Context, user *pb.User) error {
	err := w.Repo.CreateUser(ctx, w.Cfg, user)
	if err != nil {
		log.Error().Err(err).Msg("failed create user in db")
		return err
	}
	return nil
}
