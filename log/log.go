package log

import (
	"strings"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {
	logger = zerolog.New(zerolog.ConsoleWriter{
		TimeFormat: "2006-01-02 15:04:05",
	}).With().Logger()
}

func Info(template string, args ...any) {
	logger.Info().Msgf(replaceTemplate(template), args...)
}

func Error(template string, args ...any) {
	logger.Error().Msgf(replaceTemplate(template), args...)
}

func replaceTemplate(template string) string {
	return strings.ReplaceAll(template, "{}", "%v") + "\n"
}

func SetLogger(_logger zerolog.Logger) {
	logger = _logger
}
