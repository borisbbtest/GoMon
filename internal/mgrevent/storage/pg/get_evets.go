package storagepg

import (
	"context"
	"embed"

	"github.com/borisbbtest/GoMon/internal/mgrevent/models"
	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
)

var SQLFile embed.FS

func (hook *StoreDBinPostgreSQL) GetEvents(ctx context.Context, eve []*mgrevent.Event) (err error, res []*models.PGEvent) {
	SQLFile.ReadDir("../../migrations")

	return
}
