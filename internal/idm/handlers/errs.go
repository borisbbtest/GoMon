package handlers

import (
	"errors"

	"google.golang.org/grpc/codes"

	"github.com/borisbbtest/GoMon/internal/idm/service"
)

// ErrCodesMapping - функция маппинга ошибок приложения на кастомные статусы сервиса и gRPC Codes
func ErrCodesMapping(err error) (string, codes.Code) {
	switch {
	case errors.Is(err, service.ErrWrongPassword):
		return "UNAUTHENTICATED", codes.Unauthenticated
	case errors.Is(err, service.ErrEmptySQLResult):
		return "NOT_FOUND", codes.NotFound
	case errors.Is(err, service.ErrUserExists):
		return "FAILED_PRECONDITION", codes.FailedPrecondition
	default:
		return "INTERNAL", codes.Internal
	}
}
