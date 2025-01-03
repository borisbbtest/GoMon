// Package handlers описывает хэндлеры для работы с бизнес логикой в приложении
package handlers

import (
	"github.com/borisbbtest/GoMon/internal/cmdb/models"
	pb "github.com/borisbbtest/GoMon/internal/models/cmdb"
)

// GRPC - структура содержащая конфиг сервера и реализующая интерфейс gRPC сервера
type GRPC struct {
	pb.UnimplementedCmdbServer
	App *models.ConfigWrapper
}
