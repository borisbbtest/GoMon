// Package postgres реализует интерфейс Storager для работы приложения IDM с использованием СУБД Postgresql
package postgres

import (
	"context"
	"fmt"

	"github.com/borisbbtest/GoMon/internal/idm/configs"
	"github.com/borisbbtest/GoMon/internal/idm/service"
	pb "github.com/borisbbtest/GoMon/internal/models/idm"
	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var log = zerolog.New(service.LogConfig()).With().Timestamp().Caller().Logger()

// IdmRepo - обьект, реализующий интерфейс Storager.
type IdmRepo struct {
	Pool *pgxpool.Pool // Connection pool для Postgres
}

// NewIdmRepo - функция, инициализирующая подключение к Postgres
func NewIdmRepo(ctx context.Context, cfg *configs.AppConfig) (*IdmRepo, error) {
	dbpool, err := pgxpool.Connect(ctx, cfg.DBDSN)
	if err != nil {
		return nil, err
	}
	repo := &IdmRepo{dbpool}
	return repo, nil
}

// Close - функция, закрывающая подключение к Postgres
func (r *IdmRepo) Close() {
	r.Pool.Close()
}

// PGUser - структура, для сканирования объектов User из БД с учетом null значений.
type PGUser struct {
	Id        pgtype.UUID
	Login     pgtype.Text
	Firstname pgtype.Text
	Lastname  pgtype.Text
	Password  pgtype.Text
	Source    pgtype.Text
	CreatedAt pgtype.Timestamptz
}

// PGSession - структура, для сканирования объектов Session из БД с учетом null значений.
type PGSession struct {
	Id       pgtype.UUID
	Config   pgtype.JSON
	Login    pgtype.Text
	Created  pgtype.Timestamptz
	Duration pgtype.Timestamptz
}

// ConvertToPB - функция приведения пользователя из Postgres к структуре из protobuf
func (u *PGUser) ConvertToPB() *pb.User {
	user := pb.User{
		Id:        fmt.Sprintf("%x-%x-%x-%x-%x", u.Id.Bytes[0:4], u.Id.Bytes[4:6], u.Id.Bytes[6:8], u.Id.Bytes[8:10], u.Id.Bytes[10:16]),
		Login:     u.Login.String,
		Firstname: u.Firstname.String,
		Lastname:  u.Lastname.String,
		Password:  u.Password.String,
		Source:    u.Source.String,
		CreatedAt: timestamppb.New(u.CreatedAt.Time),
	}
	return &user
}

// ConvertToPB - функция приведения сессии из Postgres к структуре из protobuf
func (s *PGSession) ConvertToPB() *pb.Session {
	session := pb.Session{
		Id:       fmt.Sprintf("%x-%x-%x-%x-%x", s.Id.Bytes[0:4], s.Id.Bytes[4:6], s.Id.Bytes[6:8], s.Id.Bytes[8:10], s.Id.Bytes[10:16]),
		Config:   string(s.Config.Bytes),
		Login:    s.Login.String,
		Created:  timestamppb.New(s.Created.Time),
		Duration: timestamppb.New(s.Duration.Time),
	}
	return &session
}
