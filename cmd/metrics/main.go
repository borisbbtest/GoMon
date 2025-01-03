package main

import (
	"github.com/borisbbtest/GoMon/internal/metrics/app"
	"github.com/borisbbtest/GoMon/internal/metrics/configs"
	"github.com/borisbbtest/GoMon/internal/metrics/utils"
)

func main() {
	cfg, err := configs.GetConfig()

	ap, err := app.Init(cfg)
	if err != nil {
		utils.Log.Error().Stack().Err(err)
	}
	ap.Start()

}
