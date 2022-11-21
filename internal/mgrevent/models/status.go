package models

import "github.com/jackc/pgx/pgtype"

// Status определяет критичность события

type Status struct {
	Id   pgtype.UUID
	Name pgtype.Text
	Code pgtype.Numeric
}
