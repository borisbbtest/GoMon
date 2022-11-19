package storage

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/mgrevent/models"
	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
)

type Storage interface {
	Close()
	SaveEvents(ctx context.Context, eve []*mgrevent.Event) (err error, qerr error)
	GetEvents(ctx context.Context, eve []*mgrevent.Event) (err error, res []*models.PGEvent)
	GetSeverity(ctx context.Context, id int32) (err error, res string)
	GetStatus(ctx context.Context, id int32) (err error, res string)
}
