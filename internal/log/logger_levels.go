package log

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
	au "github.com/logrusorgru/aurora"
)

// Info show a formatted log info message.
func Info(args ...interface{}) {
	logger.Info(
		au.Green(
			emoji.Sprintf("✅  " + fmt.Sprintf("%v", args...)),
		).BgBlack().Bold().Underline(),
	)
}

// Infof show a formatted log info message.
func Infof(format string, args ...interface{}) {
	logger.Infof(
		format,
		au.Green(
			emoji.Sprintf("✅  "+format, args...),
		).BgBlack().Bold().Underline(),
	)
}

// Infoln show a formatted log info message.
func Infoln(args ...interface{}) {
	logger.Infoln(
		au.Green(
			emoji.Sprintf("✅  " + fmt.Sprintf("%v", args...)),
		).BgBlack().Bold().Underline(),
	)
}

// Debug show a formatted log debug message.
func Debug(args ...interface{}) {
	logger.Debug(
		au.Cyan(
			emoji.Sprintf("🆗  " + fmt.Sprintf("%v", args...)),
		).BgBlack().Bold().Underline(),
	)
}

// Debugf show a formatted log debug message.
func Debugf(format string, args ...interface{}) {
	logger.Debugf(
		format,
		au.Cyan(
			emoji.Sprintf("🆗  "+format, args...),
		).BgBlack().Bold().Underline(),
	)
}

// Debugln show a formatted log debug message.
func Debugln(args ...interface{}) {
	logger.Debugln(
		au.Cyan(
			emoji.Sprintf("🆗  " + fmt.Sprintf("%v", args...)),
		).BgBlack().Bold().Underline(),
	)
}

// Warn show a formatted log warn message.
func Warn(args ...interface{}) {
	logger.Warn(
		au.Yellow(
			emoji.Sprintf(":warning:  " + fmt.Sprintf("%v", args...)),
		).BgBlack().Bold().Underline(),
	)
}

// Warnf show a formatted log warn message.
func Warnf(format string, args ...interface{}) {
	logger.Warnf(
		format,
		au.Yellow(
			emoji.Sprintf(":warning:  "+format, args...),
		).BgBlack().Bold().Underline(),
	)
}

// Warnln show a formatted log warn message.
func Warnln(args ...interface{}) {
	logger.Warnln(
		au.Yellow(
			emoji.Sprintf(":warning:  " + fmt.Sprintf("%v", args...)),
		).BgBlack().Bold().Underline(),
	)
}

// Error show a formatted log error message.
func Error(args ...interface{}) {
	logger.Error(
		au.BrightRed(
			emoji.Sprintf("😡  " + fmt.Sprintf("%v", args...)),
		))
}

// Errorln show a formatted log error message.
func Errorln(args ...interface{}) {
	logger.Errorln(
		au.BrightRed(
			emoji.Sprintf("😡  " + fmt.Sprintf("%v", args...)),
		).BgBlack().Bold().Underline(),
	)
}
