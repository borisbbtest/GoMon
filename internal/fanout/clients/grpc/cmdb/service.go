package integrationcmdb

import (
	"github.com/borisbbtest/GoMon/internal/models/cmdb"
)

// ServiceWrapperCmdb - Класс по работе с клиентом CMDB
type ServiceWrapperCmdb struct {
	Ci cmdb.CmdbClient //соединений gRPC Events
}
