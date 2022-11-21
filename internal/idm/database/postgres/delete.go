package postgres

import (
	"context"
	"embed"

	"github.com/borisbbtest/GoMon/internal/idm/configs"
)

// Файлы SQL для удаления записей из таблиц хранятся в директории migrations/delete/
//
//go:embed migrations/delete/*.sql
var SQLDelete embed.FS

// DeleteUser - функция, удаляющая пользователя из БД на основании Login
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

// DeleteSession - функция, удаляющая сессию из БД на основании Login и SesssionId
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

// ClearExpiredSessions - удаляет сессии с истекшим сроком expired
func (r *IdmRepo) ClearExpiredSessions(ctx context.Context, cfg *configs.AppConfig) error {
	sqlBytes, err := SQLDelete.ReadFile("migrations/delete/SQLDeleteExpiredSession.sql")
	if err != nil {
		return err
	}
	sqlQuery := string(sqlBytes)
	_, err = r.Pool.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
	return nil
}

// TruncateTables - функция, удаляющая записи из таблиц для очистки хранилища
func (r *IdmRepo) TruncateTables(ctx context.Context, cfg *configs.AppConfig) error {
	sqlBytes, err := SQLDelete.ReadFile("migrations/delete/SQLTruncateTables.sql")
	if err != nil {
		return err
	}
	sqlQuery := string(sqlBytes)
	_, err = r.Pool.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
	return nil
}
