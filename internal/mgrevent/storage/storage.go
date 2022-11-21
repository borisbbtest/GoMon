package storage

import (
	"context"
	"time"

	"github.com/borisbbtest/GoMon/internal/mgrevent/models"
	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
)

// Интерфейс стореджа
type Storage interface {
	Close()
	SaveEvents(ctx context.Context, eve []*mgrevent.Event) (err error, qerr error)
	GetSeverity(ctx context.Context, id int32) (err error, res string)
	GetStatus(ctx context.Context, id int32) (err error, res string)

	GetEvent(ctx context.Context, eve *mgrevent.Event) (err error, res *models.Events)
	GetEventDuration(ctx context.Context, startTime time.Time, endTime time.Time) (err error, res *models.Events)
}
