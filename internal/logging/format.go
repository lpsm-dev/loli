package logging

import (
	"fmt"
	"runtime"

	"github.com/lpmatos/loli/internal/helpers"
	"github.com/sirupsen/logrus"
)

// TextFormatter function - return logrus Text formatter.
func TextFormatter(colors bool) *logrus.TextFormatter {
	return &logrus.TextFormatter{
		DisableTimestamp:       false,
		FullTimestamp:          true,
		TimestampFormat:        timestampFormat,
		DisableColors:          colors,
		DisableSorting:         false,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", helpers.FormatFilePath(f.File), f.Line)
		},
	}
}

// JSONFormatter function - return logrus JSON formatter.
func JSONFormatter(pretty bool) *logrus.JSONFormatter {
	if pretty {
		return &logrus.JSONFormatter{
			DisableTimestamp: false,
			TimestampFormat:  timestampFormat,
			PrettyPrint:      true,
		}
	}
	return &logrus.JSONFormatter{
		DisableTimestamp: false,
		TimestampFormat:  timestampFormat,
		PrettyPrint:      false,
	}
}

// PlainFormatter struct.
type PlainFormatter struct {
	TimestampFormat string
	LevelDesc       []string
}

// Format function - return logrus Plain formatter.
func (plain *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s %s %s\n", plain.LevelDesc[entry.Level], plain.TimestampFormat, entry.Message)), nil
}
