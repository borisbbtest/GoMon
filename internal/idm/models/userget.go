package models

import (
	"context"

	pb "github.com/borisbbtest/GoMon/internal/models/idm"
)

func (w *ConfigWrapper) GetUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	user, err := w.Repo.GetUser(ctx, w.Cfg, user.Login)
	if err != nil {
		log.Error().Err(err).Msg("failed get user")
		return nil, err
	}
	return user, nil
}

func (w *ConfigWrapper) GetAllUsers(ctx context.Context) ([]*pb.User, error) {
	users, err := w.Repo.GetAllUsers(ctx, w.Cfg)
	if err != nil {
		log.Error().Err(err).Msg("failed get list users")
		return nil, err
	}
	return users, nil
}
