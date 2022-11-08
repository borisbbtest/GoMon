package postgres

import (
	"fmt"

	pb "github.com/borisbbtest/GoMon/internal/idm/proto/idm"
	"github.com/jackc/pgx/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PGUser struct {
	Id        pgtype.UUID
	Login     pgtype.Text
	Firstname pgtype.Text
	Lastname  pgtype.Text
	Password  pgtype.Text
	Source    pgtype.Text
	CreatedAt pgtype.Timestamptz
}

type PGSession struct {
	Id       pgtype.UUID
	Config   pgtype.JSON
	Login    pgtype.Text
	Created  pgtype.Timestamptz
	Duration pgtype.Timestamptz
}

func (u *PGUser) ConvertToPB() *pb.User {
	user := pb.User{
		Id:        fmt.Sprintf("%x-%x-%x-%x-%x", u.Id.Bytes[0:4], u.Id.Bytes[4:6], u.Id.Bytes[6:8], u.Id.Bytes[8:10], u.Id.Bytes[10:16]),
		Login:     u.Login.String,
		Firstname: u.Firstname.String,
		Lastname:  u.Lastname.String,
		Password:  u.Password.String,
		Source:    u.Source.String,
		CreatedAt: timestamppb.New(u.CreatedAt.Time),
	}
	return &user
}

func (s *PGSession) ConvertToPB() *pb.Session {
	session := pb.Session{
		Id:       fmt.Sprintf("%x-%x-%x-%x-%x", s.Id.Bytes[0:4], s.Id.Bytes[4:6], s.Id.Bytes[6:8], s.Id.Bytes[8:10], s.Id.Bytes[10:16]),
		Config:   string(s.Config.Bytes),
		Login:    s.Login.String,
		Created:  timestamppb.New(s.Created.Time),
		Duration: timestamppb.New(s.Duration.Time),
	}
	return &session
}
