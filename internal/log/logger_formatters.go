package log

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/lpmatos/loli/internal/helpers"
	"github.com/sirupsen/logrus"
)

var (
	// Default log format will output [INFO]: 2006-01-02T15:04:05Z07:00 - Log message
	defaultLogFormat              = "[%lvl%]: %time% - %msg%\n"
	defaultTimestampFormat string = "2006-01-02 15:04:05"
)

func textFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		DisableColors:          true,
		TimestampFormat:        defaultTimestampFormat,
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
		TimestampFormat:           defaultTimestampFormat,
	}
}

func jsonFormatter(pretty bool) *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		DisableTimestamp: false,
		TimestampFormat:  defaultTimestampFormat,
		PrettyPrint:      pretty,
	}
}

// PlainFormatter implements logrus.Formatter interface.
type PlainFormatter struct {
	// Timestamp format
	TimestampFormat string
	// Available standard keys: time, msg, lvl
	// Also can include custom fields but limited to strings.
	// All of fields need to be wrapped inside %% i.e %time% %msg%
	LogFormat string
}

// Format function - return logrus custom Plain formatter.
func (plain *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := plain.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := plain.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	output = strings.Replace(output, "%msg%", entry.Message, 1)

	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%lvl%", level, 1)

	for k, val := range entry.Data {
		switch v := val.(type) {
		case string:
			output = strings.Replace(output, "%"+k+"%", v, 1)
		case int:
			s := strconv.Itoa(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		case bool:
			s := strconv.FormatBool(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}

	return []byte(output), nil
}
