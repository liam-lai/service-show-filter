package util

import (
	"time"

	"github.com/rs/zerolog"
)

func LogInit(LogLevel string) {
	level, err := zerolog.ParseLevel(LogLevel)
	if err != nil {
		panic(err)
	}
	zerolog.SetGlobalLevel(level)

	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}
}
