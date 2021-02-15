package log

import (
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

// Logrus implements Logger interface.
type Logrus struct {
	logger *logrus.Logger
}

// NewLogger is a delegator method for logrus.New.
func NewLogger() *logrus.Logger {
	return logrus.New()
}

func applyLoggerOptions(opts []Option) *options {
	var o options

	o.logger = logger

	for _, opt := range opts {
		opt(&o)
	}

	return &o
}

// Setup returns a new logrus instance.
func Setup(opts ...Option) error {
	conf := applyLoggerOptions(opts)
	conf.logger.SetFormatter(conf.logger.Formatter)
	conf.logger.SetLevel(conf.logger.Level)
	return nil
}

func init() {
	logger = NewLogger()
}
