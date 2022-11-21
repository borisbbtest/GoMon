package integrationcmdb

import (
	"github.com/borisbbtest/GoMon/internal/models/cmdb"
)

type ServiceWrapperCmdb struct {
	Ci cmdb.CmdbClient //соединений gRPC Events
}
