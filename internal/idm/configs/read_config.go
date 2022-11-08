package configs

import (
	"flag"
	"strconv"

	"github.com/caarlos0/env"
)

func (cfg *AppConfig) envRead() {
	err := env.Parse(cfg)
	if err != nil {
		log.Error().Err(err).Msg("problem with environment read")
	}
}

func (cfg *AppConfig) flagsRead() {
	flag.Func("d", "key for database DSN, example: -d \"postgres://pi:toor@192.168.1.69:5432/yandex\"", func(flagValue string) error {
		if flagValue != "" {
			cfg.DBDSN = flagValue
		}
		return nil
	})
	flag.Func("g", "server gRPC address like <server>:<port>, example: -a \"127.0.0.1:8080\"", func(flagValue string) error {
		if flagValue != "" {
			cfg.ServerAddressGRPC = flagValue
		}
		return nil
	})
	flag.Func("r", "true/false for restore metrics from disk after restart, example: -r=true", func(flagValue string) error {
		if flagValue != "" {
			value, err := strconv.ParseBool(flagValue)
			if err != nil {
				return err
			}
			cfg.ReInit = value
		}
		return nil
	})
	flag.Parse()
}
