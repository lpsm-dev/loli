package log

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
	au "github.com/logrusorgru/aurora"
)

// Info function to print a pretty formatted log info message.
func Info(args ...interface{}) {
	logger.Info(
		au.Green(
			emoji.Sprintf("✅  " + fmt.Sprintf("%v", args...)),
		).Bold().Underline(),
	)
}

// Infof function to print a pretty formatted log info message.
func Infof(format string, args ...interface{}) {
	logger.Infof(
		format,
		au.Green(
			emoji.Sprintf("✅  "+format, args...),
		).Bold().Underline(),
	)
}

// Infoln function to print a pretty formatted log info message.
func Infoln(args ...interface{}) {
	logger.Infoln(
		au.Green(
			emoji.Sprintf("✅  " + fmt.Sprintf("%v\n", args...)),
		).Bold().Underline(),
	)
}

// Debug function to print a pretty formatted log debug message.
func Debug(args ...interface{}) {
	logger.Debug(
		au.Cyan(
			emoji.Sprintf("🆗  " + fmt.Sprintf("%v", args...)),
		).Bold().Underline(),
	)
}

// Debugf function to print a pretty formatted log debug message.
func Debugf(format string, args ...interface{}) {
	logger.Debugf(
		format,
		au.Cyan(
			emoji.Sprintf("🆗  "+format, args...),
		).Bold().Underline(),
	)
}

// Debugln function to print a pretty formatted log debug message.
func Debugln(args ...interface{}) {
	logger.Debugln(
		au.Cyan(
			emoji.Sprintf("🆗  " + fmt.Sprintf("%v\n", args...)),
		).Bold().Underline(),
	)
}

// Warn function to print a pretty formatted log warn message.
func Warn(args ...interface{}) {
	logger.Warn(
		au.Yellow(
			emoji.Sprintf("😲  " + fmt.Sprintf("%v", args...)),
		).Bold().Underline(),
	)
}

// Warnf function to print a pretty formatted log warn message.
func Warnf(format string, args ...interface{}) {
	logger.Warnf(
		format,
		au.Yellow(
			emoji.Sprintf("😲  "+format, args...),
		).Bold().Underline(),
	)
}

// Warnln function to print a pretty formatted log warn message.
func Warnln(args ...interface{}) {
	logger.Warnln(
		au.Yellow(
			emoji.Sprintf("😲  " + fmt.Sprintf("%v\n", args...)),
		).Bold().Underline(),
	)
}

// Error function to print a pretty formatted log error message.
func Error(args ...interface{}) {
	logger.Error(
		au.BrightRed(
			emoji.Sprintf("😡  " + fmt.Sprintf("%v", args...)),
		))
}

// Errorf function to print a pretty formatted log error message.
func Errorf(format string, args ...interface{}) {
	logger.Errorf(
		format,
		au.BrightRed(
			emoji.Sprintf("😡  "+format, args...),
		).Bold().Underline(),
	)
}

// Errorln function to print a pretty formatted log error message.
func Errorln(args ...interface{}) {
	logger.Errorln(
		au.BrightRed(
			emoji.Sprintf("😡  " + fmt.Sprintf("%v\n", args...)),
		).Bold().Underline(),
	)
}
