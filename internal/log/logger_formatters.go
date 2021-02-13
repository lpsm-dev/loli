package log

import (
	"fmt"
	"runtime"

	"github.com/lpmatos/loli/internal/helpers"
	"github.com/sirupsen/logrus"
)

var (
	timestampFormat string   = "2006-01-02 15:04:05"
	customLevels    []string = []string{
		"😋 INFO",
		"🥳 DEBUG",
		"😠 WARN",
		"😡 PANIC",
		"🤬 FATAL",
		"😡 ERROR",
	}
)

func textFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		DisableColors:          true,
		TimestampFormat:        timestampFormat,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", helpers.FormatFilePath(f.File), f.Line)
		},
	}
}

func colorFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           timestampFormat,
	}
}

func jsonFormatter(pretty bool) *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		DisableTimestamp: false,
		TimestampFormat:  timestampFormat,
		PrettyPrint:      pretty,
	}
}

type plainFormatter struct {
	TimestampFormat string
	LevelDesc       []string
}

// Format function - return logrus custom Plain formatter.
func (plain *plainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s %s %s\n",
			plain.LevelDesc[entry.Level],
			plain.TimestampFormat,
			entry.Message)),
		nil
}
