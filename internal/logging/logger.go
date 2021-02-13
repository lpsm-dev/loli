package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	logger                   = logrus.StandardLogger()
	timestampFormat string   = "2006-01-02 15:04:05"
	levels          []string = []string{
		"😡 PANIC",
		"🤬 FATAL",
		"😡 ERROR",
		"😠 WARN",
		"😋 INFO",
		"🥳 DEBUG",
	}
)

func init() {
	// Enfore our own defaults on the logrus stdlogger
	logger.Out = os.Stderr
	logger.Formatter = &logrus.TextFormatter{}
	logger.Level = logrus.InfoLevel
	logger.SetReportCaller(true)
}
