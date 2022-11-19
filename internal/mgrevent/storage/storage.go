package storage

import "github.com/borisbbtest/GoMon/internal/mgrevent/models"

type Storage interface {
	Close()
	SaveEvent(eve []models.Event) (err error)
	GetEvents(eve []models.Event) (err error, res []models.Event)
	GetSeverity(int32) (err error, res string)
	GetStatus(int32) (err error, res string)
}
