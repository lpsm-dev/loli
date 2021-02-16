package log

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// Logrus implements Logger interface.
type Logrus struct {
	logger *logrus.Logger
}

// newLogger is a delegator method for logrus.New.
func newLogger() *logrus.Logger {
	return logrus.New()
}
