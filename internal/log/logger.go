package log

import "github.com/sirupsen/logrus"

// NewLogger is a delegator method for logrus.New.
func NewLogger() *logrus.Logger {
	return logrus.New()
}
