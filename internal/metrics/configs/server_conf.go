package configs

import (
	"context"
	goflag "flag"
	"io/ioutil"

	"github.com/borisbbtest/GoMon/internal/metrics/utils"
	"github.com/caarlos0/env"
	"gopkg.in/yaml.v2"

	flag "github.com/spf13/pflag"
)

type MainConfig struct {
	DatabaseURI string `yaml:"DATABASE_URI" env:"DATABASE_URI"`
	RunAddress  string `yaml:"RUN_ADDRESS_RPC" env:"RUN_ADDRESS_RPC"`
	Ctx         context.Context
}
type ServerConfig interface {
	GetConfig() (config *MainConfig, err error)
}

func GetConfig() (config *MainConfig, err error) {
	config = &MainConfig{}
	var configFileName string

	config.Ctx = context.Background()

	utils.Log.Info().Msgf("context", "system_loyalty")
	flag.StringVarP(&config.DatabaseURI, "database_uri", "d", config.DatabaseURI, "Base URL")
	flag.StringVarP(&config.RunAddress, "run_server", "a", config.RunAddress, "Run server")
	flag.StringVarP(&configFileName, "config", "c", "", "path to the configuration file")
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	err = env.Parse(config)
	if err != nil {
		utils.Log.Error().Discard().Msgf("can't start the listening thread: %s", err)
		return
	}

	configFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		utils.Log.Error().Msgf("can't open the config file: %s", err)

	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		utils.Log.Error().Msgf("YAML can't read the config file: %s", err)
	}
	//***postgres:5432/praktikum?sslmode=disable
	utils.Log.Info().Msgf("Configuration loaded")
	return
}
