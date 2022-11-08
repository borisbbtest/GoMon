package postgres

import "errors"

var (
	ErrUserExists     = errors.New("user already exist")
	ErrEmptySQLResult = errors.New("no result in sql")
)
