package models

import (
	"encoding/json"
	"time"

	"github.com/borisbbtest/GoMon/internal/fanout/utils"
	pb "github.com/borisbbtest/GoMon/internal/models/idm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Session - внутренний тип сессии для данного модуля, используется при Unmarshall из входных данных HTTP
type Session struct {
	Id       string    `json:"id"`               // айди сессии (заполняется системой), ключевой атрибут
	Config   string    `json:"config,omitempty"` // параметр для хранения конфигураций конкретной сессии
	Login    string    `json:"login"`            // имя пользователя, ключевой атрибут
	Duration time.Time `json:"duration"`         // дата, когда сессия перестанет быть валидной (заполняется системой)
	Created  time.Time `json:"created"`          // дата создания сессии (заполняется системой)
	Code     string    `json:"code,omitempty"`   // код??
}

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
		utils.Log.Error().Err(err).Msg("failed unmarshall json")
		return err
	}
	if AliasValue.Created != "" {
		res, err := time.Parse(time.RFC3339, AliasValue.Created)
		if err != nil {
			utils.Log.Error().Err(err).Msg("failed unmarshall Created")
			return err
		}
		s.Created = res
	}
	if AliasValue.Duration != "" {
		res, err := time.Parse(time.RFC3339, AliasValue.Duration)
		if err != nil {
			utils.Log.Error().Err(err).Msg("failed unmarshall Duration")
			return err
		}
		s.Duration = res
	}
	return nil
}
