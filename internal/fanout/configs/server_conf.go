// Package пакет по созданию конфигурации для приложения
package config

import (
	goflag "flag"
	"io/ioutil"

	"github.com/borisbbtest/GoMon/internal/fanout/utils"
	"github.com/caarlos0/env"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

// MainConfig - конфиг приложения
type MainConfig struct {
	AccrualSystemAddress string `yaml:"ACCRUAL_SYSTEM_ADDRESS" env:"ACCRUAL_SYSTEM_ADDRESS"` // Адрес
	DatabaseURI          string `yaml:"DATABASE_URI" env:"DATABASE_URI"`                     // Подключение к БД
	RunAddress           string `yaml:"RUN_ADDRESS" env:"RUN_ADDRESS"`                       // ИП и порт на котором прослушивается http приложение
	EnableHTTPS          bool   `yaml:"ENABLE_HTTPS" env:"ENABLE_HTTPS"`                     // Запускаем сервер в tls
	TrustedSubnet        string `yaml:"TRUSTED_SUBNET" env:"TRUSTED_SUBNET"`                 // Ограничиваем доступ по IP
	ServiceGRpcEvents    string `yaml:"CLIENT_ADDRESS_EVENT" env:"CLIENT_ADDRESS_EVENT"`     // Адрес подключения к менеджеру событий
	ServiceGRpcMetric    string `yaml:"CLIENT_ADDRESS_METRIC" env:"CLIENT_ADDRESS_METRIC"`   // Адрес подключения к менеджеру метрик
	ServiceGRpcIDM       string `yaml:"CLIENT_ADDRESS_IDM" env:"CLIENT_ADDRESS_IDM"`         // Адрес подключения к менеджеру IDM
	ServiceGRpcCMD       string `yaml:"CLIENT_ADDRESS_CMD" env:"CLIENT_ADDRESS_CMD"`         // Адрес подключения к менеджеру CMD
}

// ServerConfig интерфейс который должен соблюдаться
type ServerConfig interface {
	GetConfig() (config *MainConfig, err error)
}

// GetConfig - чтение конфигурации из файла переменных флагов
func GetConfig() (config *MainConfig, err error) {
	config = &MainConfig{}
	var configFileName string
	flag.StringVarP(&config.AccrualSystemAddress, "accrual_system_adders", "r", config.AccrualSystemAddress, "Accrual system address")
	flag.StringVarP(&config.DatabaseURI, "database_uri", "d", config.DatabaseURI, "Base URL")
	flag.StringVarP(&config.RunAddress, "run_server", "a", config.RunAddress, "Run server")
	flag.BoolVarP(&config.EnableHTTPS, "tls", "s", false, "In HTTP server is Enable TLS")
	flag.StringVarP(&config.TrustedSubnet, "trusted_subnet", "t", config.RunAddress, "Trust subnet")
	flag.StringVarP(&configFileName, "config", "c", "./config/fanout.yaml", "path to the configuration file")
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
