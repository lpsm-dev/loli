package log

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type nopCloser struct{}

func (nopCloser) Close() error { return nil }

func applyLoggerOptions(opts []LoggerOption) *loggerConfig {
	logConf := loggerConfig{
		logger:    logger,
		level:     logrus.InfoLevel,
		formatter: colorFormatter(),
		writer:    os.Stdout,
	}

	for _, value := range opts {
		value(&logConf)
	}

	return &logConf
}

// Initialize function - initializes the logrus logger.
func Initialize(opts ...LoggerOption) error {
	conf := applyLoggerOptions(opts)

	conf.logger.SetFormatter(conf.formatter)
	conf.logger.SetLevel(conf.level)

	dirPath, _ := filepath.Abs(filepath.Dir(conf.outputPath))
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.Mkdir(dirPath, 0775)
	}

	file, err := os.OpenFile(conf.outputPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	conf.logger.SetOutput(file)

	// Only output the warnings _after_ the logger has been configured
	for _, warning := range conf.warnings {
		conf.logger.Warn(warning)
	}

	return nil
}
