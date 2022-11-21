package models

import (
	"encoding/json"
	"time"

	pb "github.com/borisbbtest/GoMon/internal/models/cmdb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToPB - конвертация типа Ci в protobuf Ci
func (c *Ci) ToPB() *pb.Ci {
	return &pb.Ci{
		Name:        c.Name,
		Description: c.Description,
		Update:      timestamppb.New(c.Update),
		Created:     timestamppb.New(c.Created),
		CreatedBy:   c.CreatedBy,
		Type:        c.Type,
	}
}

// UnmarshalJSON - функция переопределяющия правила анмаршалера для timestamp в Ci
func (c *Ci) UnmarshalJSON(data []byte) error {
	type CiAlias Ci
	AliasValue := &struct {
		*CiAlias
		Created string `json:"created,omitempty"`
		Update  string `json:"update,omitempty"`
	}{
		CiAlias: (*CiAlias)(c),
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
		c.Created = res
	}
	if AliasValue.Update != "" {
		res, err := time.Parse(time.RFC3339, AliasValue.Update)
		if err != nil {
			log.Error().Err(err).Msg("failed unmarshall Update")
			return err
		}
		c.Update = res
	}
	return nil
}
