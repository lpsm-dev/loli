package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/lpmatos/loli/internal/helpers"
	"github.com/sirupsen/logrus"
)

var (
	timestampFormat string   = "2006-01-02 15:04:05"
	levels          []string = []string{"PANC", "FATL", "ERRO", "WARN", "INFO", "DEBG"}
	logFile                  = "/var/log/loli.log"
)

// PlainFormatter struct.
type PlainFormatter struct {
	TimestampFormat string
	LevelDesc       []string
}

// Format function - return logrus Plain formatter.
func (plain *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s %s %s\n", plain.LevelDesc[entry.Level], plain.TimestampFormat, entry.Message)), nil
}

// InitLog function - initializes the logrus logger.
func InitLog(logLevel, logFormatter, logOutput string) error {
	logrus.SetReportCaller(true)

	// 1 - Logrus Set Formatter
	switch strings.ToLower(logFormatter) {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			DisableTimestamp: false,
			TimestampFormat:  timestampFormat,
			PrettyPrint:      false,
		})
	case "json-pretty":
		logrus.SetFormatter(&logrus.JSONFormatter{
			DisableTimestamp: false,
			TimestampFormat:  timestampFormat,
			PrettyPrint:      true,
		})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableTimestamp:       false,
			FullTimestamp:          true,
			TimestampFormat:        timestampFormat,
			DisableSorting:         false,
			DisableLevelTruncation: true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				return "", fmt.Sprintf("%s:%d", helpers.FormatFilePath(f.File), f.Line)
			},
		})
	case "plain":
		plainFormatter := new(PlainFormatter)
		plainFormatter.TimestampFormat = timestampFormat
		plainFormatter.LevelDesc = levels
		logrus.SetFormatter(plainFormatter)
	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:            true,
			DisableTimestamp:       false,
			FullTimestamp:          true,
			TimestampFormat:        timestampFormat,
			DisableSorting:         false,
			DisableLevelTruncation: true,
		})
	}

	// 2 - Logrus Set OutPut
	switch strings.ToLower(logOutput) {
	case "stderr":
		logrus.SetOutput(os.Stderr)
	case "stdout":
		logrus.SetOutput(os.Stdout)
	case "file":
		dirPath, _ := filepath.Abs(filepath.Dir(logFile))
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			os.Mkdir(dirPath, 0775)
		}
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return err
		}
		logrus.SetOutput(file)
	default:
		logrus.Warn("No log output defined, using stderr")
		logrus.SetOutput(os.Stderr)
	}

	// 3 - Logrus Parse Level
	level, error := logrus.ParseLevel(logLevel)
	if error != nil {
		logrus.SetLevel(logrus.DebugLevel)
		return error
	}
	logrus.SetLevel(level)

	return nil
}
