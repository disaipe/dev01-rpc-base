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
	fileWriter := &lumberjack.Logger{
		Filename: path.Join(Config.GetWorkingDir(), "app.log"),
		MaxSize:  20,
		MaxAge:   90,
	}

	writers := []io.Writer{fileWriter}

	if !Config.IsService() {
		consoleWriter := &zerolog.ConsoleWriter{Out: os.Stderr}
		writers = append(writers, consoleWriter)
	}

	zerolog.TimeFieldFormat = time.DateTime
	Logger = zerolog.New(io.MultiWriter(writers...)).With().Timestamp().Logger()

	err := fileWriter.Rotate()
	if err != nil {
		Logger.Warn().Msgf("Failed to rotate logs: %v", err)
	}
}
