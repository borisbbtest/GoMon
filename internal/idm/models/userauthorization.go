package models

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	pb "github.com/borisbbtest/GoMon/internal/idm/proto/idm"
	"github.com/borisbbtest/GoMon/internal/idm/service"
)

func (w *ConfigWrapper) AuthorizeUser(ctx context.Context, login string, password string) (*pb.Session, error) {
	user, err := w.GetUser(ctx, &pb.User{Login: login})
	if err != nil {
		log.Error().Err(err).Msg("failed get user")
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, service.ErrWrongPassword
	}
	session, err := w.CreateSession(ctx, user)
	if err != nil {
		log.Error().Err(err).Msg("failed create session in db")
		return nil, err
	}
	return session, nil
}
