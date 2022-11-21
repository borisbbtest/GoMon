package models

import (
	integrationmgrevent "github.com/borisbbtest/GoMon/internal/fanout/clients/grpc/events"
	integrationidm "github.com/borisbbtest/GoMon/internal/fanout/clients/grpc/idm"
	integrationmetric "github.com/borisbbtest/GoMon/internal/fanout/clients/grpc/metrics"
	config "github.com/borisbbtest/GoMon/internal/fanout/configs"
)

type ClientPool struct {
	Idm *integrationidm.ServiceWrapperIdm // клиент для модуля idm (управление пользователями)
	// Cmdb    cmdb.CmdbClient                  // клиент для модуля cmdb (управление КЕ)
	Metrics  *integrationmetric.ServiceWrapperMetric   // клиент для модуля metrics (хранение метрик)
	Mgrevent *integrationmgrevent.ServiceWrapperEvents // клиент для модуля events (хранение событий)
}

// NewPoolService - создаем пул конектов
func NewPoolService(cfg config.MainConfig) (res *ClientPool) {
	res = &ClientPool{
		Idm:      integrationidm.NewConnIdm(cfg),
		Metrics:  integrationmetric.NewConnMetric(cfg),
		Mgrevent: integrationmgrevent.NewConnMgrEvent(cfg),
	}
	return
}

// FanInContextKey - тип, использующийся при создании context withValue для избежания возможных коллизий
type FanInContextKey string
