package postgres

import (
	"context"
	"embed"

	"github.com/borisbbtest/GoMon/internal/idm/configs"
)

//go:embed migrations/delete/*.sql
var SQLDelete embed.FS

func (r *IdmRepo) DeleteUser(ctx context.Context, cfg *configs.AppConfig, login string) error {
	sqlBytes, err := SQLDelete.ReadFile("migrations/delete/SQLDeleteUser.sql")
	if err != nil {
		return err
	}
	sqlQuery := string(sqlBytes)
	_, err = r.Pool.Exec(ctx, sqlQuery, login)
	if err != nil {
		return err
	}
	return nil
}

func (r *IdmRepo) DeleteSession(ctx context.Context, cfg *configs.AppConfig, login string, id string) error {
	sqlBytes, err := SQLDelete.ReadFile("migrations/delete/SQLDeleteSession.sql")
	if err != nil {
		return err
	}
	sqlQuery := string(sqlBytes)
	_, err = r.Pool.Exec(ctx, sqlQuery, login, id)
	if err != nil {
		return err
	}
	return nil
}
