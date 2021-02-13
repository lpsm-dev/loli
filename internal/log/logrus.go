package log

import (
	"github.com/kyokomi/emoji/v2"
	"github.com/logrusorgru/aurora"
)

// Create a new instance of the logger.
var logger = NewLogger()

// Infof logs a formatted info message.
func Infof(format string, args ...interface{}) {
	logger.Info(aurora.Cyan(emoji.Sprintf(":beer: "+format, args...)))
}

// Debugf logs a formatted debug message.
func Debugf(format string, args ...interface{}) {
	logger.Debug(aurora.Green(emoji.Sprintf(":pizza:"+format, args...)))
}

// Successf logs a formatted success message.
func Successf(format string, a ...interface{}) {
	logger.Info(aurora.Green(emoji.Sprintf(":white_check_mark: "+format, a...)))
}

// Warnf logs a formatted warn message.
func Warnf(format string, args ...interface{}) {
	logger.Warn(aurora.Yellow(emoji.Sprintf(":exclamation: "+format, args...)))
}

// Errorf logs a formatted error message.
func Errorf(format string, args ...interface{}) {
	logger.Error(aurora.Red(emoji.Sprintf(":exclamation: "+format, args...)))
}

// Fatalln logs a formatted fatal message.
func Fatalln(format string, args ...interface{}) {
	logger.Fatalln(aurora.Red(emoji.Sprintf(":exclamation: "+format, args...)))
}
