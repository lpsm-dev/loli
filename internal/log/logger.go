package log

import "github.com/sirupsen/logrus"

var (
	logger *logrus.Logger
)

type logrusLogger struct {
	logger *logrus.Logger
}

// NewLogger is a delegator method for logrus.New.
func NewLogger() *logrus.Logger {
	return logrus.New()
}

func init() {
	logger = NewLogger()
}
