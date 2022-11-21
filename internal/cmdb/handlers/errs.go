package handlers

import (
	"errors"

	"google.golang.org/grpc/codes"

	"github.com/borisbbtest/GoMon/internal/cmdb/service"
)

// ErrCodesMapping - функция маппинга ошибок приложения на кастомные статусы сервиса и gRPC Codes
func ErrCodesMapping(err error) (string, codes.Code) {
	switch {
	case errors.Is(err, service.ErrInsertObjects), errors.Is(err, service.ErrSelectObjects), errors.Is(err, service.ErrDeleteObjects):
		return "DATA_LOSS", codes.DataLoss
	case errors.Is(err, service.ErrEmptySQLResult):
		return "NOT_FOUND", codes.NotFound
	case errors.Is(err, service.ErrObjectExists):
		return "FAILED_PRECONDITION", codes.FailedPrecondition
	default:
		return "INTERNAL", codes.Internal
	}
}
