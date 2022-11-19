package storagepg

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/mgrevent/utils"
	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
)

func (hook *StoreDBinPostgreSQL) SaveEvents(ctx context.Context, eve []*mgrevent.Event) (err error, qerr error) {

	conn, err := hook.pgp.NewConn()
	if err != nil {
		utils.Log.Error().Msgf("Save Events ", err)
		return err, nil
	}
	tx, err := conn.PostgresPool.Begin(ctx)

	query, err := SQLFileInit.ReadFile("migrations/core/insert_events.sql")

	if err != nil {
		utils.Log.Error().Msgf("Save Events ", err)
		return err, nil
	}

	b := &pgx.Batch{}
	for _, v := range eve {

		b.Queue(string(query),
			v.Title,
			v.Description,
			v.Source,
			v.Status,
			v.Created.AsTime(),
			v.Update.AsTime(),
			v.Key,
			v.KeyClose,
			pq.Array(v.Assigned),
			v.Severity,
			v.AutoRunner,
			pq.Array(v.RelarionCi))
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
