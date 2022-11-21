// Package  пакет по созданию стордже на основе СУБД Postgres
package storagepg

import (
	"context"
	"embed"

	"github.com/borisbbtest/GoMon/internal/metrics/configs"
	"github.com/borisbbtest/GoMon/internal/metrics/database/postgres"
	"github.com/borisbbtest/GoMon/internal/metrics/utils"
)

type StoreDBinPostgreSQL struct {
	pgp     postgres.Plugin
	connStr string
}

//go:embed migrations/*/*.sql
var SQLFileInit embed.FS

// Создает БД
func (hook *StoreDBinPostgreSQL) ExecSqlInitFiles(ctx context.Context, sqlfiles []string) (err error) {

	var query []byte

	conn, err := hook.pgp.NewConn()

	if err != nil {
		utils.Log.Error().Msgf("Init Files ", err)
		return err
	}

	for _, q := range sqlfiles {

		if query, err = SQLFileInit.ReadFile(q); err != nil {
			utils.Log.Debug().Msgf(err.Error())
			return err
		}

		if _, err := conn.PostgresPool.Exec(ctx, string(query)); err != nil {
			utils.Log.Debug().Msgf(err.Error())
			//return err
		}
	}
	return
}

// Конструктор стореджа в СУБД Postgres
func NewPostgreSQLStorage(connStr *configs.MainConfig) (res *StoreDBinPostgreSQL, err error) {
	res = &StoreDBinPostgreSQL{}
	res.connStr = connStr.DatabaseURI
	res.pgp.Start(connStr.DatabaseURI)

	filesSql := []string{
		"migrations/init/ext.sql",
		"migrations/init/schema.sql",
		"migrations/init/metrics.sql",
		"migrations/init/alter.sql"}

	res.ExecSqlInitFiles(connStr.Ctx, filesSql)
	return
}
func (hook *StoreDBinPostgreSQL) Close() {
	hook.pgp.Stop()
}
