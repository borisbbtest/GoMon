package main

import (
	"github.com/borisbbtest/GoMon/internal/fanout/app"
	config "github.com/borisbbtest/GoMon/internal/fanout/configs"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
)

func main() {
	cfg, err := config.GetConfig()

	ap, err := app.Init(cfg)
	if err != nil {
		utils.Log.Error().Stack().Err(err)
	}
	ap.Start()

}
