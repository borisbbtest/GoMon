package models

import (
	"encoding/json"
	"time"

	pb "github.com/borisbbtest/GoMon/internal/models/mgrevent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// textToSeverity - парсинг string в protobuf Severity
// func textToSeverity(severity string) (pb.Severity, error) {
// 	if v, ok := pb.Severity_value[severity]; ok {
// 		pbsev := pb.Severity(v)
// 		return pbsev, nil
// 	}
// 	intsev, err := strconv.ParseInt(severity, 10, 32)
// 	if err != nil {
// 		return pb.Severity(0), err
// 	}
// 	if _, ok := pb.Severity_name[int32(intsev)]; ok {
// 		pbsev := pb.Severity(int32(intsev))
// 		return pbsev, nil
// 	}
// 	return pb.Severity(0), service.ErrEventWrongSeverity
// }

// textToStatus - парсинг string в protobuf Status
// func textToStatus(status string) (pb.Status, error) {
// 	if v, ok := pb.Status_value[status]; ok {
// 		pbstatus := pb.Status(v)
// 		return pbstatus, nil
// 	}
// 	intstatus, err := strconv.ParseInt(status, 10, 32)
// 	if err != nil {
// 		return pb.Status(0), err
// 	}
// 	if _, ok := pb.Status_name[int32(intstatus)]; ok {
// 		pbstatus := pb.Status(int32(intstatus))
// 		return pbstatus, nil
// 	}
// 	return pb.Status(0), service.ErrEventWrongStatus
// }

// ToPB - конвертация типа Ci в protobuf  Ci.
// Возвращает ошибку в случае, когда Severity или Status не имеют совпадений с protobuf
func (e *Event) ToPB() *pb.Event {
	return &pb.Event{
		Title:       e.Title,
		Description: e.Description,
		Source:      e.Source,
		Status:      e.Status,
		Created:     timestamppb.New(e.Created),
		Update:      timestamppb.New(e.Update),
		Key:         e.Key,
		KeyClose:    e.KeyClose,
		Assigned:    e.Assigned,
		AutoRunner:  e.AutoRunner,
		Severity:    e.Severity,
		RelarionCi:  e.RelarionCi,
		CreatedBy:   e.CreatedBy,
		Count:       e.Count,
		Uuid:        e.Uuid,
	}
}

// UnmarshalJSON - функция переопределяющия правила анмаршалера для timestamp в Event
func (e *Event) UnmarshalJSON(data []byte) error {
	type EventAlias Event
	AliasValue := &struct {
		*EventAlias
		Created string `json:"created"`
		Update  string `json:"update"`
	}{
		EventAlias: (*EventAlias)(e),
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
		e.Created = res
	}
	if AliasValue.Update != "" {
		res, err := time.Parse(time.RFC3339, AliasValue.Update)
		if err != nil {
			log.Error().Err(err).Msg("failed unmarshall Update")
			return err
		}
		e.Update = res
	}
	return nil
}
