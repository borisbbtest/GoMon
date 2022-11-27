// Package по работе с сервисом IDM
package integrationidm

import (
	"context"
	"time"

	pb "github.com/borisbbtest/GoMon/internal/models/idm"
)

// RegisterUser - метод, регистрирующий user в idm с конвертацией User в protobuf User. Сразу же создается сессия для этого пользователя.
func (hook *ServiceWrapperIdm) RegisterUser(ctx context.Context, user *User) (Session, error) {
	var req pb.RegistrationRequest
	req.Person = user.ToPB()
	resp, err := hook.Idm.Registration(ctx, &req)
	if err != nil {
		return Session{}, err
	}
	var session Session
	session.FromPB(resp.Ss)
	return session, nil
}

// AuthorizeUser - метод, авторизирующий user в системе с использованием idm. Сразу же создается сессия для этого пользователя.
func (hook *ServiceWrapperIdm) AuthorizeUser(ctx context.Context, user *User) (Session, error) {
	var req pb.AuthorizationRequest
	req.Login = user.Login
	req.Password = user.Password
	resp, err := hook.Idm.Authorization(ctx, &req)
	if err != nil {
		return Session{}, err
	}
	var session Session
	session.FromPB(resp.Ss)
	return session, nil
}

// CheckAuthorized - метод, возвращают true в случае, если сессия существует в idm и время ее окончания после текущего времени.
// Во всех остальных случаях возвращается false
func (hook *ServiceWrapperIdm) CheckAuthorized(ctx context.Context, login string, id string) bool {
	var req pb.GetSessionRequest
	req.Id = id
	req.Login = login
	resp, err := hook.Idm.GetSession(ctx, &req)
	if err != nil {
		return false
	}
	ss := resp.Ss
	return time.Now().Before(ss.Duration.AsTime())
}
