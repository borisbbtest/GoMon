package models

import (
	"context"
	"time"

	pb "github.com/borisbbtest/GoMon/internal/models/idm"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (w *ConfigWrapper) CreateSession(ctx context.Context, user *pb.User) (*pb.Session, error) {
	SessionId := uuid.New().String()
	session := &pb.Session{
		Id:       SessionId,
		Login:    user.Login,
		Created:  timestamppb.Now(),
		Duration: timestamppb.New(time.Now().Add(w.Cfg.SessionTimeExpired)),
	}
	err := w.Repo.CreateSession(ctx, w.Cfg, session)
	if err != nil {
		log.Error().Err(err).Msg("failed create session in db")
		return nil, err
	}
	return session, nil
}
