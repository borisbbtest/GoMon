package postgres

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/idm/configs"
	"github.com/borisbbtest/GoMon/internal/idm/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)

var log = zerolog.New(service.LogConfig()).With().Timestamp().Caller().Logger()

type IdmRepo struct {
	Pool *pgxpool.Pool
}

func NewIdmRepo(ctx context.Context, cfg *configs.AppConfig) (*IdmRepo, error) {
	dbpool, err := pgxpool.Connect(ctx, cfg.DBDSN)
	if err != nil {
		return nil, err
	}
	repo := &IdmRepo{dbpool}
	return repo, nil
}
