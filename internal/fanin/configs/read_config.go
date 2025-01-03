package configs

import (
	"flag"
	"os"

	"github.com/caarlos0/env"
	"gopkg.in/yaml.v3"
)

// YamlRead - функция считывания конфига из файла yaml
func (cfg *AppConfig) YamlRead(file string) error {
	yfile, err := os.ReadFile(file)
	if err != nil {
		log.Error().Err(err).Msg("file open trouble")
		return err
	} else {
		err = yaml.Unmarshal(yfile, &cfg)
		if err != nil {
			log.Error().Err(err).Msg("parse yaml err")
			return err
		}
	}
	return nil
}

// EnvRead - функция считывания конфига из переменных окружения
func (cfg *AppConfig) EnvRead() error {
	err := env.Parse(cfg)
	if err != nil {
		log.Error().Err(err).Msg("problem with environment read")
		return err
	}
	return nil
}

// FlagsRead - функция считывания конфига из флагов запуска
func (cfg *AppConfig) FlagsRead() {
	flag.Func("s", "self address like <server>:<port>, example: -a \"127.0.0.1:8080\"", func(flagValue string) error {
		if flagValue != "" {
			cfg.HTTPServerAddress = flagValue
		}
		return nil
	})
	flag.Func("i", "service idm address like <server>:<port>, example: -a \"127.0.0.1:8080\"", func(flagValue string) error {
		if flagValue != "" {
			cfg.IdmAddress = flagValue
		}
		return nil
	})
	flag.Func("c", "service cmdb address like <server>:<port>, example: -a \"127.0.0.1:8080\"", func(flagValue string) error {
		if flagValue != "" {
			cfg.CmdbAddress = flagValue
		}
		return nil
	})
	flag.Func("m", "service metrics address like <server>:<port>, example: -a \"127.0.0.1:8080\"", func(flagValue string) error {
		if flagValue != "" {
			cfg.MetricsAddress = flagValue
		}
		return nil
	})
	flag.Func("e", "service events address like <server>:<port>, example: -a \"127.0.0.1:8080\"", func(flagValue string) error {
		if flagValue != "" {
			cfg.EventsAddress = flagValue
		}
		return nil
	})
	flag.Parse()
}
