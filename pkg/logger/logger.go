package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

type LevelWriter struct {
	Console bool
}

func (LevelWriter) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

func (w LevelWriter) WriteLevel(level zerolog.Level, b []byte) (int, error) {
	out := os.Stdout
	if level >= zerolog.ErrorLevel {
		out = os.Stderr
	}
	if w.Console {
		return zerolog.ConsoleWriter{Out: out}.Write(b)
	}
	return out.Write(b)
}

func App() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(LevelWriter{Console: true}).
		With().
		Timestamp().
		Caller().
		Logger()
	log.Logger = logger
}
