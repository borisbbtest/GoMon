// Package менеджер подключений к БД postgres
package postgres

import (
	"time"

	"github.com/borisbbtest/GoMon/internal/mgrevent/utils"
)

// Plugin класс объекта менеджера подключения
type Plugin struct {
	connMgr *connManager
	dsn     string
}

type requestHandler func(conn *postgresConn, key string, params []interface{}) (res interface{}, err error)

// Start запускает иницилизацию к БД
func (p *Plugin) Start(dsn string) {
	p.dsn = dsn
	p.connMgr = p.NewConnManager(
		time.Duration(20000)*time.Second,
		time.Duration(20000)*time.Second,
	)
}

// Sop останавливает коннект к БД
func (p *Plugin) Stop() {
	p.connMgr.stop()
	p.connMgr = nil
}

// NewConn - создает создает отдельное содинение к БД
func (p *Plugin) NewConn() (conn *postgresConn, err error) {

	conn, err = p.connMgr.GetPostgresConnection(p.dsn)
	if err != nil {
		utils.Log.Info().Msgf("connection error: %s", err)
		return nil, err
	}
	return
}
