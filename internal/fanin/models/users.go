package models

import (
	"context"

	pb "github.com/borisbbtest/GoMon/internal/models/idm"
)

// RegisterUser - метод, регистрирующий user в idm с конвертацией User в protobuf User. Сразу же создается сессия для этого пользователя.
func (cw *ConfigWrapper) RegisterUser(ctx context.Context, user *User) (Session, error) {
	var req pb.RegistrationRequest
	req.Person = user.ToPB()
	resp, err := cw.Conns.Idm.Registration(ctx, &req)
	if err != nil {
		return Session{}, err
	}
	var session Session
	session.FromPB(resp.Ss)
	return session, nil
}

// AuthorizeUser - метод, авторизирующий user в системе с использованием idm. Сразу же создается сессия для этого пользователя.
func (cw *ConfigWrapper) AuthorizeUser(ctx context.Context, user *User) (Session, error) {
	var req pb.AuthorizationRequest
	req.Login = user.Login
	req.Password = user.Password
	resp, err := cw.Conns.Idm.Authorization(ctx, &req)
	if err != nil {
		return Session{}, err
	}
	var session Session
	session.FromPB(resp.Ss)
	return session, nil
}
