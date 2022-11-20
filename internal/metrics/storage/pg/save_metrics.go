package storagepg

import (
	"context"
	"time"

	"github.com/borisbbtest/GoMon/internal/metrics/utils"
	"github.com/borisbbtest/GoMon/internal/models/metrics"
	"github.com/jackc/pgx/v4"
)

// Записываем  массив данных в БД
func (hook *StoreDBinPostgreSQL) SaveEvents(ctx context.Context, eve []*metrics.Metric) (err error, qerr error) {

	conn, err := hook.pgp.NewConn()
	if err != nil {
		utils.Log.Error().Msgf("Save Metrics ", err)
		return err, nil
	}
	tx, err := conn.PostgresPool.Begin(ctx)

	query, err := SQLFileInit.ReadFile("migrations/core/insert_metrics.sql")

	if err != nil {
		utils.Log.Error().Msgf("Save Metrics ", err)
		return err, nil
	}

	b := &pgx.Batch{}
	for _, v := range eve {

		b.Queue(string(query),
			v.Name,
			v.Value,
			time.Now(),
			v.SourceFromSystems,
			v.RelationCi,
			v.SourceTime.AsTime(),
			metrics.Types(v.Tp),
		)
	}

	batchResults := tx.SendBatch(ctx, b)

	var rows pgx.Rows
	for qerr == nil {
		rows, qerr = batchResults.Query()
		if qerr != nil && qerr.Error() != "no result" {
			utils.Log.Debug().Msgf("Add %s", qerr)
		}
		rows.Close()
	}
	tx.Commit(ctx)

	return
}
