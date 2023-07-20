package rpc

import (
	"io"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger zerolog.Logger

func init() {
	consoleWriter := &zerolog.ConsoleWriter{Out: os.Stderr}
	fileWriter := &lumberjack.Logger{
		Filename: path.Join(Config.GetWorkingDir(), "app.log"),
		MaxSize:  20,
		MaxAge:   90,
	}

	writers := []io.Writer{consoleWriter, fileWriter}

	zerolog.TimeFieldFormat = time.DateTime
	Logger = zerolog.New(io.MultiWriter(writers...)).With().Timestamp().Logger()

	err := fileWriter.Rotate()
	if err != nil {
		Logger.Warn().Msgf("Failed to rotate logs: %v", err)
	}
}
