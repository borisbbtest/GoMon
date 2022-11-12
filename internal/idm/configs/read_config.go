package configs

import (
	"flag"
	"os"
	"strconv"
	"time"

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
	flag.Func("t", "time duration for session live, example: -t \"100s\"", func(flagValue string) error {
		if flagValue != "" {
			interval, err := time.ParseDuration(flagValue)
			if err != nil {
				return err
			}
			cfg.SessionTimeExpired = interval
		}
		return nil
	})
	flag.Parse()
}
