package models

import (
	"encoding/json"
	"time"

	pb "github.com/borisbbtest/GoMon/internal/models/idm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToPB - конвертация типа User в protobuf User
func (u *User) ToPB() *pb.User {
	return &pb.User{
		Login:     u.Login,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Password:  u.Password,
		CreatedAt: timestamppb.New(u.CreatedAt),
		Source:    u.Source,
		Id:        u.Id,
	}
}

// UnmarshalJSON - функция переопределяющия правила анмаршалера для timestamp в User
func (u *User) UnmarshalJSON(data []byte) error {
	type UserAlias User
	AliasValue := &struct {
		*UserAlias
		CreatedAt string `json:"CreatedAt"`
	}{
		UserAlias: (*UserAlias)(u),
	}
	if err := json.Unmarshal(data, AliasValue); err != nil {
		log.Error().Err(err).Msg("failed unmarshall json")
		return err
	}
	if AliasValue.CreatedAt != "" {
		res, err := time.Parse(time.RFC3339, AliasValue.CreatedAt)
		if err != nil {
			log.Error().Err(err).Msg("failed unmarshall CreatedAt")
			return err
		}
		u.CreatedAt = res
	}
	return nil
}
