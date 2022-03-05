package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	once   sync.Once
	logger zerolog.Logger
)

func GetLogger() zerolog.Logger {
	once.Do(func() {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	})

	return logger
}
