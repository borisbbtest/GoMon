package models

import (
	"encoding/json"

	integrationcmdb "github.com/borisbbtest/GoMon/internal/fanout/clients/grpc/cmdb"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
)

// RequestGetCis структура  для запроса HTTP для массива КЕ
type RequestGetCis struct {
	Item *[]string
}

// ResponseGetCis  структура  для ответа HTTP для Массива КЕ
type ResponseGetCis struct {
	Root *[]*integrationcmdb.Ci
}

// ResponseGetCi  структура  для ответа HTTP для одного КЕ
type ResponseGetCi struct {
	Root *integrationcmdb.Ci
}

// UnmarshalJSON - функция переопределяющия правила анмаршалера для timestamp в Metric
func (hook *RequestGetCis) ParseRequest(data []byte) error {
	Req := &struct {
		ListCisName []string `json:"ListCisName"`
	}{}
	if err := json.Unmarshal(data, Req); err != nil {
		utils.Log.Error().Err(err).Msg("failed unmarshall json")
		return err
	}
	hook.Item = &Req.ListCisName
	return nil
}
