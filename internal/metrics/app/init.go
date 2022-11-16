package app

import (
	"github.com/borisbbtest/GoMon/internal/metrics/utils"
)

type ServiceEvents struct {
}

func (hook *ServiceEvents) Start() (err error) {
	utils.Log.Error().Msg("Hello ")
	return nil
}
