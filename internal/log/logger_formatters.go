package log

import (
	"github.com/sirupsen/logrus"
)

var defaultTimestampFormat string = "2006-01-02 15:04:05"

// Configures the format of the log output to use "text" formatter.
func textFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		DisableColors:             true,
		ForceColors:               false,
		EnvironmentOverrideColors: false,
		FullTimestamp:             true,
		TimestampFormat:           defaultTimestampFormat,
		DisableLevelTruncation:    true,
	}
}

// Configures the format of the log output to use "color" formatter.
func colorFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		DisableColors:             false,
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		TimestampFormat:           defaultTimestampFormat,
		DisableLevelTruncation:    true,
	}
}

// Configures the format of the log output to use "json" formatter.
func jsonFormatter(pretty bool) *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		TimestampFormat:  defaultTimestampFormat,
		DisableTimestamp: false,
		PrettyPrint:      pretty,
	}
}
