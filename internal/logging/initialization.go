package logging

import (
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/client9/reopen"
	"github.com/sirupsen/logrus"
)

type nopCloser struct{}

func (nopCloser) Close() error { return nil }

func applyLoggerOptions(opts []LoggerOption) *loggerConfig {
	conf := loggerConfig{
		logger:    logger,
		level:     logrus.InfoLevel,
		formatter: &logrus.TextFormatter{},
		writer:    os.Stdout,
	}

	for _, v := range opts {
		v(&conf)
	}
	return &conf
}

// Initialize function - initializes the logrus logger.
func Initialize(opts ...LoggerOption) (io.Closer, error) {
	conf := applyLoggerOptions(opts)

	// Being unable to open the output file will cause an error
	writer, closer, err := getOutputWriter(conf)
	if err != nil {
		return closer, err
	}

	conf.logger.SetFormatter(conf.formatter)
	conf.logger.SetLevel(conf.level)
	conf.logger.SetOutput(writer)

	// Only output the warnings _after_ the logger has been configured
	for _, warning := range conf.warnings {
		conf.logger.Warn(warning)
	}

	return closer, nil
}

func getOutputWriter(conf *loggerConfig) (io.Writer, io.Closer, error) {
	if conf.writer != nil {
		return conf.writer, nopCloser{}, nil
	}

	// When writing to a file, use `reopen` so that we can
	// reopen the file on SIGHUP signals
	f, err := reopen.NewFileWriterMode(conf.outputPath, 0644)
	if err != nil {
		return f, nopCloser{}, err
	}

	isMainLogger := conf.logger == logger

	sighup := make(chan os.Signal, 1)
	signal.Notify(sighup, syscall.SIGHUP)
	go listenForSignalHangup(f, isMainLogger, conf.outputPath, sighup)

	return f, f, nil
}

// Will listen for SIGHUP signals and reopen the underlying file.
func listenForSignalHangup(l reopen.Reopener, isMainLogger bool, logFilePath string, sighup chan os.Signal) {
	for v := range sighup {
		// Specifically, do _not_ write to the log that is being reopened,
		// but log this to the _main_ log file instead as the actual log
		// might be specialised, eg: an access logger leading to an incorrect entry
		logger.WithFields(logrus.Fields{"signal": v, "path": logFilePath}).Print("Reopening log file on signal")

		err := l.Reopen()
		if err != nil {
			if isMainLogger {
				// Main logger failed to reopen, last ditch effort to notify the user, but don't
				// do this for auxiliary loggers, since we may get double-logs
				log.Printf("Unable to reopen log file '%s' after %v. Error %v", logFilePath, v, err)
			} else {
				logger.WithError(err).WithFields(logrus.Fields{"signal": v, "path": logFilePath}).Print("Failed to reopen log file")
			}
		}
	}
}
