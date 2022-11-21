package postgres

import (
	"context"
	"embed"

	"github.com/borisbbtest/GoMon/internal/cmdb/configs"
)

// Файлы SQL для удаления записей из таблиц хранятся в директории migrations/delete/
//
//go:embed migrations/delete/*.sql
var SQLDelete embed.FS

// DeleteUser - функция, удаляющая КЕ из БД на основании Login
func (r *CmdbRepo) DeleteObject(ctx context.Context, cfg *configs.AppConfig, name string) error {
	sqlBytes, err := SQLDelete.ReadFile("migrations/delete/SQLDeleteCi.sql")
	if err != nil {
		return err
	}
	sqlQuery := string(sqlBytes)
	_, err = r.Pool.Exec(ctx, sqlQuery, name)
	if err != nil {
		return err
	}
	return nil
}

// TruncateTables - функция, удаляющая записи из таблиц для очистки хранилища
func (r *CmdbRepo) TruncateTables(ctx context.Context, cfg *configs.AppConfig) error {
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
