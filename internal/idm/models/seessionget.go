package models

import (
	"context"

	pb "github.com/borisbbtest/GoMon/internal/idm/proto/idm"
)

func (w *AppWrapper) GetSession(ctx context.Context, login string, id string) (*pb.Session, error) {
	session, err := w.Repo.GetSession(ctx, w.Cfg, login, id)
	if err != nil {
		log.Error().Err(err).Msg("failed get user")
		return nil, err
	}
	return session, nil
}

func (w *AppWrapper) GetAllSessions(ctx context.Context) ([]*pb.Session, error) {
	sessions, err := w.Repo.GetAllSessions(ctx, w.Cfg)
	if err != nil {
		log.Error().Err(err).Msg("failed get list sessions")
		return nil, err
	}
	return sessions, nil
}
