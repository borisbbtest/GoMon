package storage

import model "github.com/borisbbtest/GoMon/internal/fanout/models"

type SessionHTTP struct {
	DBSession map[string]model.DataUser
}
type Storage interface {
}
