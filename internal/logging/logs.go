package logging

import (
	"strings"

	logrus "github.com/sirupsen/logrus"
)

// logger variable - local variable logrus logger.
var logger *logrus.Logger

// Setup function - configure logrus logger.
func Setup() {
	logger = logrus.New()
	logger.SetReportCaller(false)
	Formatter := &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		DisableColors:   false,
	}
	logger.SetFormatter(Formatter)
}

// SetLogLevel sets the logrus logging level.
func SetLogLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.Fatalf("Unknown level: %s", level)
	}
}
