package storagepg

import (
	"embed"

	"github.com/borisbbtest/GoMon/internal/mgrevent/models"
)

var SQLFile embed.FS

func (hook *StoreDBinPostgreSQL) GetEvents(eve []models.Event) (err error, res []models.Event) {
	SQLFile.ReadDir("../../migrations")

	return
}
