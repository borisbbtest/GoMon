package configs

import (
	goflag "flag"

	"github.com/borisbbtest/GoMon/internal/metrics/utils"
	"github.com/caarlos0/env"

	flag "github.com/spf13/pflag"
)

type MainConfig struct {
	AccrualSystemAddress string `yaml:"ACCRUAL_SYSTEM_ADDRESS" env:"ACCRUAL_SYSTEM_ADDRESS"`
	DatabaseURI          string `yaml:"DATABASE_URI" env:"DATABASE_URI"`
	RunAddress           string `yaml:"RUN_ADDRESS" env:"RUN_ADDRESS"`
}
type ServerConfig interface {
	GetConfig() (config *MainConfig, err error)
}

func GetConfig() (config *MainConfig, err error) {
	config = &MainConfig{}

	utils.Log.Info().Msgf("context", "system_loyalty")
	flag.StringVarP(&config.AccrualSystemAddress, "accrual_system_adders", "r", config.AccrualSystemAddress, "Accrual system address")
	flag.StringVarP(&config.DatabaseURI, "database_uri", "d", config.DatabaseURI, "Base URL")
	flag.StringVarP(&config.RunAddress, "run_server", "a", config.RunAddress, "Run server")
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	err = env.Parse(config)
	if err != nil {
		utils.Log.Error().Discard().Msgf("can't start the listening thread: %s", err)
		return
	}

	//***postgres:5432/praktikum?sslmode=disable
	utils.Log.Info().Msgf("Configuration loaded")
	return
}
