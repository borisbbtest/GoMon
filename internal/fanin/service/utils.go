// Package service содержит в себе служебные инструменты для внутренней работы приложения
package service

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// LogConfig - функция конфигурации логгера zerolog
func LogConfig() zerolog.ConsoleWriter {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	return output
}
