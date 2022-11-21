package postgres

import (
	"context"
	"embed"
	"time"

	"github.com/jackc/pgconn"

	"github.com/borisbbtest/GoMon/internal/cmdb/configs"
	"github.com/borisbbtest/GoMon/internal/cmdb/service"
	pb "github.com/borisbbtest/GoMon/internal/models/cmdb"
)

// Файлы SQL для вставки записей в таблицы хранятся в директории migrations/insert/
//
//go:embed migrations/insert/*.sql
var SQLInsert embed.FS

// CreateObject - функция создает КЕ, полученное по gRPC
func (r *CmdbRepo) CreateObject(ctx context.Context, cfg *configs.AppConfig, ci *pb.Ci) error {
	sqlBytes, err := SQLInsert.ReadFile("migrations/insert/SQLInsertNewCi.sql")
	if err != nil {
		return err
	}
	sqlQuery := string(sqlBytes)
	time := time.Now()
	_, err = r.Pool.Exec(ctx, sqlQuery, ci.Name, ci.Description, time, time, ci.CreatedBy, ci.Type)
	if err != nil {
		pgerr := err.(*pgconn.PgError)
		if pgerr.Code == "23505" {
			return service.ErrObjectExists
		}
		return err
	}
	return nil
}
