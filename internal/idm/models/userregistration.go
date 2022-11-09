package models

import (
	"context"

	pb "github.com/borisbbtest/GoMon/internal/idm/proto/idm"
)

func (w *ConfigWrapper) RegisterUser(ctx context.Context, user *pb.User) (*pb.Session, error) {
	err := w.CreateUser(ctx, user)
	if err != nil {
		log.Error().Err(err).Msg("failed create user in db")
		return nil, err
	}
	session, err := w.CreateSession(ctx, user)
	if err != nil {
		log.Error().Err(err).Msg("failed create session in db")
		return nil, err
	}
	return session, nil
}
