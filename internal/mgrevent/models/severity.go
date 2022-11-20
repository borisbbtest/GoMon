package models

import (
	"github.com/jackc/pgx/pgtype"
)

type Severity struct {
	id   pgtype.UUID
	name pgtype.Text
	code pgtype.Numeric
}

func SeverityList() (res Severity) {
	return
}

func (hook *Severity) GetSeverity(x int) (res int) {
	return
}
