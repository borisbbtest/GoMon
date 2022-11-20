package models

import "github.com/jackc/pgx/pgtype"

type Status struct {
	Id   pgtype.UUID
	Name pgtype.Text
	Code pgtype.Numeric
}
