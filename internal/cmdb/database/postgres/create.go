package postgres

import (
	"context"
	"embed"

	"github.com/borisbbtest/GoMon/internal/cmdb/configs"
)

// Файлы SQL для создания таблиц хранятся в директории migrations/create/
//
//go:embed migrations/create/*.sql
var SQLCreate embed.FS

// CreateTables - метод создают таблицы при инициализации подключения, если они отсутствуют.
func (r *CmdbRepo) CreateTables(ctx context.Context, cfg *configs.AppConfig) error {
	sqlBytes, err := SQLCreate.ReadFile("migrations/create/SQLCreateCiTable.sql")
	if err != nil {
		return err
	}
	sqlQuery := string(sqlBytes)
	ct, err := r.Pool.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
	log.Info().Str("initialize table", ct.String())
	return nil
}
