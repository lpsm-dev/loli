package log

import (
	"github.com/lpmatos/loli/internal/constants"
	"github.com/sirupsen/logrus"
)

// Configure the logrus format to use "text" formatter.
func textFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		DisableColors:             true,
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		TimestampFormat:           constants.DefaultTimestampFormat,
		DisableLevelTruncation:    true,
	}
}

// Configure the logrus format to use "color" formatter.
func colorFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		DisableColors:             false,
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		TimestampFormat:           constants.DefaultTimestampFormat,
		DisableLevelTruncation:    true,
	}
}

// Configure the logrus format to use "json" formatter.
func jsonFormatter(pretty bool) *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		TimestampFormat:  constants.DefaultTimestampFormat,
		DisableTimestamp: false,
		PrettyPrint:      pretty,
	}
}
