package loggers

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once

var log zerolog.Logger

func Get() zerolog.Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		logLevel := int(zerolog.InfoLevel)

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		if os.Getenv("APP_ENV") != "development" {
			fileLogger := &lumberjack.Logger{
				Filename:   "logs/app.log",
				MaxSize:    5, //
				MaxBackups: 10,
				MaxAge:     14,
				Compress:   true,
			}

			output = zerolog.MultiLevelWriter(os.Stderr, fileLogger)
		}

		// var gitRevision string

		// buildInfo, ok := debug.ReadBuildInfo()

		// if ok {
		// 	for _, v := range buildInfo.Settings {
		// 		if v.Key == "vcs.revision" {
		// 			gitRevision = v.Value
		// 			break
		// 		}
		// 	}
		// }

		log = zerolog.New(output).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			// Str("git_revision", gitRevision).
			// Str("go_version", buildInfo.GoVersion).
			Logger()
	})

	return log
}

func GetError() zerolog.Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		logLevel := int(zerolog.ErrorLevel)

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		if os.Getenv("APP_ENV") != "development" {
			fileLogger := &lumberjack.Logger{
				Filename:   "logs/app.log",
				MaxSize:    5, //
				MaxBackups: 10,
				MaxAge:     14,
				Compress:   true,
			}

			output = zerolog.MultiLevelWriter(os.Stderr, fileLogger)
		}

		// var gitRevision string

		// buildInfo, ok := debug.ReadBuildInfo()

		// if ok {
		// 	for _, v := range buildInfo.Settings {
		// 		if v.Key == "vcs.revision" {
		// 			gitRevision = v.Value
		// 			break
		// 		}
		// 	}
		// }

		log = zerolog.New(output).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			// Str("git_revision", gitRevision).
			// Str("go_version", buildInfo.GoVersion).
			Logger()
	})

	return log
}
