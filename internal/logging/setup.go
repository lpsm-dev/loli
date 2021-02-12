package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	defaultLogLevel          = logrus.DebugLevel
	timestampFormat string   = "2006-01-02 15:04:05"
	levels          []string = []string{"PANC", "FATL", "ERRO", "WARN", "INFO", "DEBG"}
	outputFile      *os.File
)

// Setup function - configure logrus logger.
func Setup(logTarget, logFile, logFormatter, logLevel string, colors bool) {
	logrus.SetReportCaller(true)
	setLogLevel(logLevel)
	setLogFormat(logFormatter, colors)
	switch logTarget {
	case "stdout":
		logrus.SetOutput(os.Stdout)
	case "stderr":
		logrus.SetOutput(os.Stderr)
	case "file":
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err == nil {
			logrus.SetOutput(file)
		} else {
			logrus.Errorf("Open file <%s> for logging failed - <%v>! Setup defaulting output to stdout\n", file, err)
			logrus.SetOutput(os.Stdout)
		}
	default:
		logrus.Errorln("Error: log_target invalid, defaulting to Stdout")
		logrus.SetOutput(os.Stdout)
	}
}
