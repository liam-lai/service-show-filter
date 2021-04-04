package middleware

import (
	"os"

	"github.com/rs/zerolog"
)

func Logger() *zerolog.Logger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &logger
}
