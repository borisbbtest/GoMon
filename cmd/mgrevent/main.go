package main

import (
	"github.com/borisbbtest/GoMon/internal/mgrevent/app"
	"github.com/borisbbtest/GoMon/internal/mgrevent/configs"
	"github.com/borisbbtest/GoMon/internal/mgrevent/utils"
)

func main() {
	cfg, err := configs.GetConfig()

	ap, err := app.Init(cfg)
	if err != nil {
		utils.Log.Error().Stack().Err(err)
	}
	ap.Start()

}
