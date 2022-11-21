package models

import (
	"encoding/json"
	"time"

	pb "github.com/borisbbtest/GoMon/internal/models/idm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToPB - конвертация типа Session в protobuf Session
func (s *Session) ToPB() *pb.Session {
	return &pb.Session{
		Id:       s.Id,
		Config:   s.Config,
		Login:    s.Login,
		Duration: timestamppb.New(s.Duration),
		Created:  timestamppb.New(s.Created),
		Code:     s.Code,
	}
}

// FromPB - конвертация типа protobuf Session в Session
func (s *Session) FromPB(session *pb.Session) {
	s.Id = session.Id
	s.Config = session.Config
	s.Login = session.Login
	s.Duration = session.Duration.AsTime()
	s.Created = session.Created.AsTime()
	s.Code = session.Code
}

// UnmarshalJSON - функция переопределяющия правила анмаршалера для timestamp в Session
func (s *Session) UnmarshalJSON(data []byte) error {
	type SessionAlias Session
	AliasValue := &struct {
		*SessionAlias
		Created  string `json:"created"`
		Duration string `json:"duration"`
	}{
		SessionAlias: (*SessionAlias)(s),
	}
	if err := json.Unmarshal(data, AliasValue); err != nil {
		log.Error().Err(err).Msg("failed unmarshall json")
		return err
	}
	if AliasValue.Created != "" {
		res, err := time.Parse(time.RFC3339, AliasValue.Created)
		if err != nil {
			log.Error().Err(err).Msg("failed unmarshall Created")
			return err
		}
		s.Created = res
	}
	if AliasValue.Duration != "" {
		res, err := time.Parse(time.RFC3339, AliasValue.Duration)
		if err != nil {
			log.Error().Err(err).Msg("failed unmarshall Duration")
			return err
		}
		s.Duration = res
	}
	return nil
}
