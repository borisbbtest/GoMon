// Package handlers описывает хэндлеры для работы с бизнес логикой в приложении
package handlers

import (
	"github.com/borisbbtest/GoMon/internal/idm/models"
	pb "github.com/borisbbtest/GoMon/internal/models/idm"
)

// GRPC - структура содержащая конфиг сервера и реализующая интерфейс gRPC сервера
type GRPC struct {
	pb.UnimplementedIdmServer
	App *models.ConfigWrapper
}
