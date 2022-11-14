package postgres

import (
	"context"
	"embed"
	"errors"

	"github.com/borisbbtest/GoMon/internal/cmdb/configs"
	"github.com/borisbbtest/GoMon/internal/cmdb/service"
	pb "github.com/borisbbtest/GoMon/internal/models/cmdb"
	"github.com/jackc/pgx/v4"
)

// Файлы SQL для вставки записей в таблицы хранятся в директории migrations/select/
//
//go:embed migrations/select/*.sql
var SQLSelect embed.FS

// GetUser - функция, которая возвращает пользователя в структуре protobuf c указанным Login
func (r *CmdbRepo) GetObject(ctx context.Context, cfg *configs.AppConfig, name string) (*pb.Ci, error) {
	sqlBytes, err := SQLSelect.ReadFile("migrations/select/SQLSelectCi.sql")
	if err != nil {
		return nil, err
	}
	sqlQuery := string(sqlBytes)
	row := r.Pool.QueryRow(ctx, sqlQuery, name)
	var ci PGCi
	err = row.Scan(&ci.Name, &ci.Description, &ci.Update, &ci.Created, &ci.CreatedBy, &ci.Type)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, service.ErrEmptySQLResult
		}
		log.Error().Err(err).Msg("failed get ci")
		return nil, err
	}
	result := ci.ConvertToPB()
	return result, nil
}
