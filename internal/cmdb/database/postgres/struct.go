// Package postgres реализует интерфейс Storager для работы приложения IDM с использованием СУБД Postgresql
package postgres

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/cmdb/configs"
	"github.com/borisbbtest/GoMon/internal/cmdb/service"
	pb "github.com/borisbbtest/GoMon/internal/models/cmdb"
	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var log = zerolog.New(service.LogConfig()).With().Timestamp().Caller().Logger()

// CmdbRepo - обьект, реализующий интерфейс Storager.
type CmdbRepo struct {
	Pool *pgxpool.Pool // Connection pool для Postgres
}

// NewCmdbRepo - функция, инициализирующая подключение к Postgres
func NewCmdbRepo(ctx context.Context, cfg *configs.AppConfig) (*CmdbRepo, error) {
	dbpool, err := pgxpool.Connect(ctx, cfg.DBDSN)
	if err != nil {
		return nil, err
	}
	repo := &CmdbRepo{dbpool}
	return repo, nil
}

// Close - функция, закрывающая подключение к Postgres
func (r *CmdbRepo) Close() {
	r.Pool.Close()
}

// PGCi - структура, для сканирования объектов КЕ из БД с учетом null значений.
type PGCi struct {
	Name        pgtype.Text
	Description pgtype.Text
	Update      pgtype.Timestamptz
	Created     pgtype.Timestamptz
	CreatedBy   pgtype.Text
	Type        pgtype.Text
}

// ConvertToPB - функция приведения пользователя из Postgres к структуре из protobuf
func (u *PGCi) ConvertToPB() *pb.Ci {
	ci := pb.Ci{
		Name:        u.Name.String,
		Description: u.Description.String,
		Update:      timestamppb.New(u.Update.Time),
		Created:     timestamppb.New(u.Created.Time),
		CreatedBy:   u.CreatedBy.String,
		Type:        u.Type.String,
	}
	return &ci
}
