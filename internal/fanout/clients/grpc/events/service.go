package integrationevents

import (
	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
)

// ServiceWrapperEvents Класс по работе с сервисом Event manager
type ServiceWrapperEvents struct {
	Events mgrevent.EventsClient //соединений gRPC  евент менеджером
}
