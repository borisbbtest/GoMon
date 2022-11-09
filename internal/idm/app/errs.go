package app

import (
	"errors"

	"google.golang.org/grpc/codes"

	"github.com/borisbbtest/GoMon/internal/idm/service"
)

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
