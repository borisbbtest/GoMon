package models

import (
	"github.com/jackc/pgx/pgtype"
)

type PGEvent struct {
	Id          pgtype.UUID
	Title       pgtype.Text
	Description pgtype.Text
	Source      pgtype.Text
	Status      pgtype.Numeric
	Created     pgtype.Timestamp
	Update      pgtype.Timestamp
	Key         pgtype.Text
	KeyClose    pgtype.Text
	Assigned    pgtype.TextArray
	Severity    pgtype.Numeric
	AutoRunner  pgtype.Text
	RelarionCi  pgtype.TextArray
}
