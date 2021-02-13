package log

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type loggerConfig struct {
	logger     *logrus.Logger
	level      logrus.Level
	formatter  logrus.Formatter
	writer     io.Writer
	outputPath string

	// A list of warnings that will be emitted once the logger is configured
	warnings []string
}

// LoggerOption will configure a new logrus Logger.
type LoggerOption func(*loggerConfig)

// WithFormatter allows setting the format to `text`, `color`, `json`, `json-pretty` or `plain`. In case
// the input is not recognized it defaults to text with a warning.
// More details of these formats:
// * `text` 				- human readable.
// * `color` 				- human readable, in color. Useful for development.
// * `json` 				- computer readable, new-line delimited JSON.
// * `json-pretty` 	- computer readable pretty JSON.
// * `plain` 				- custom human readable.
func WithFormatter(format string) LoggerOption {
	return func(conf *loggerConfig) {
		switch strings.ToLower(format) {
		case "text":
			conf.formatter = textFormatter()
		case "color":
			conf.formatter = colorFormatter()
		case "json":
			conf.formatter = jsonFormatter(false)
		case "json-pretty":
			conf.formatter = jsonFormatter(true)
		case "plain":
			plainFormatter := new(plainFormatter)
			plainFormatter.TimestampFormat = timestampFormat
			plainFormatter.LevelDesc = customLevels
			conf.formatter = plainFormatter
		default:
			conf.warnings = append(conf.warnings, fmt.Sprintf("Unknown logging format %s, ignoring option", format))
		}
	}
}

// WithLogLevel is used to set the log level when defaulting to `info` is not
// wanted. Other options are: `debug`, `warn`, `error`, `fatal`, and `panic`.
func WithLogLevel(level string) LoggerOption {
	return func(conf *loggerConfig) {
		logrusLevel, err := logrus.ParseLevel(level)
		if err != nil {
			conf.warnings = append(conf.warnings, fmt.Sprintf("Unknown log level, ignoring option: %v", err))
		} else {
			conf.level = logrusLevel
		}
	}
}

// WithOutput allows customization of the sink of the logger. Output is either:
// `stdout`, `stderr`, or `file`.
func WithOutput(output, file string) LoggerOption {
	return func(conf *loggerConfig) {
		switch strings.ToLower(output) {
		case "stdout":
			conf.writer = os.Stdout
		case "stderr":
			conf.writer = os.Stderr
		case "file":
			conf.writer = nil
			conf.outputPath = file
		default:
			conf.warnings = append(conf.warnings, fmt.Sprintf("Unknown log output, ignoring option: %v", output))
		}
	}
}
