package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/lpmatos/loli/internal/helpers"
	"github.com/sirupsen/logrus"
)

type options struct {
	logger *logrus.Logger
}

// Option will configure a new logrus Logger.
type Option func(*options) error

// WithConfig takes the logger configuration and applies it.
func WithConfig(cfg Config) Option {
	var opts []Option

	opts = append(opts, WithLogLevel(cfg.Level, cfg.Silence))
	opts = append(opts, WithFormatter(cfg.Format))
	opts = append(opts, WithOutputStr(cfg.Output, cfg.File))
	opts = append(opts, WithSetReportCaller(cfg.Details))

	// return all options
	return func(c *options) error {
		for _, opt := range opts {
			if err := opt(c); err != nil {
				return err
			}
		}
		return nil
	}
}

// WithFormatter allows setting the format to `text`, `color`, `json`, `json-pretty` or `plain`. In case
// the input is not recognized it defaults to text with a warning.
// More details of these formats:
// * `text` 				- human readable.
// * `color` 				- human readable, in color. Useful for development.
// * `json` 				- computer readable, new-line delimited JSON.
// * `json-pretty` 	- computer readable pretty JSON.
// * `plain` 				- custom human readable.
func WithFormatter(format string) Option {
	return func(o *options) error {
		switch strings.ToLower(format) {
		case "text":
			o.logger.Formatter = textFormatter()
		case "color":
			o.logger.Formatter = colorFormatter()
		case "json":
			o.logger.Formatter = jsonFormatter(false)
		case "json-pretty":
			o.logger.Formatter = jsonFormatter(true)
		default:
			o.logger.Formatter = colorFormatter()
		}
		return nil
	}
}

// WithLogLevel is used to set the log level when defaulting to `info` is not
// wanted. Other options are: `debug`, `warn`, `error`, `fatal`, and `panic`.
func WithLogLevel(level string, silence bool) Option {
	return func(opt *options) error {
		logrusLevel, err := logrus.ParseLevel(level)
		if err != nil {
			return fmt.Errorf("failed to convert level: %w", err)
		}
		if silence {
			opt.logger.Level = logrus.ErrorLevel
		} else {
			opt.logger.Level = logrusLevel
		}
		return nil
	}
}

// WithOutputStr allows customization of the sink of the logger. Output is either:
// `stdout`, `stderr`, or `file`.
func WithOutputStr(output, file string) Option {
	opt := func(*options) error { return nil }

	if output == "" {
		return opt
	}

	switch strings.ToLower(output) {
	case "stdout":
		opt = WithOutput(os.Stdout)
	case "stderr":
		opt = WithOutput(os.Stdout)
	case "file":
		helpers.MakeDirIfNotExist(file)

		f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return func(c *options) error {
				return fmt.Errorf("unable to open/create file %q: %w", output, err)
			}
		}
		opt = WithOutput(f)

		runtime.SetFinalizer(f, func(ff *os.File) {
			_ = ff.Sync()
			_ = ff.Close()
		})
	default:
		helpers.MakeDirIfNotExist(file)

		f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return func(c *options) error {
				return fmt.Errorf("unable to open/create file %q: %w", output, err)
			}
		}
		opt = WithOutput(f)

		runtime.SetFinalizer(f, func(ff *os.File) {
			_ = ff.Sync()
			_ = ff.Close()
		})
	}

	return opt
}

// WithOutput configures the writer used to write logs to.
func WithOutput(writer io.Writer) Option {
	return func(opt *options) error {
		opt.logger.Out = writer
		return nil
	}
}

// WithSetReportCaller configures the logrus logger SetReportCaller.
func WithSetReportCaller(enable bool) Option {
	return func(opt *options) error {
		opt.logger.SetReportCaller(enable)
		return nil
	}
}
