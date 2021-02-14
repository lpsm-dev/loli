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

// Warning show a formatted log warning message.
func Warning(args ...interface{}) {
	logger.Warning(
		au.BrightYellow(
			emoji.Sprintf(":warning:  " + fmt.Sprintf("%v", args...)),
		).BgBlack().Bold().Underline(),
	)
}

// Warningf show a formatted log warning message.
func Warningf(format string, args ...interface{}) {
	logger.Warningf(
		format,
		au.BrightYellow(
			emoji.Sprintf(":warning:  "+format, args...),
		).BgBlack().Bold().Underline(),
	)
}

// Warningln show a formatted log warning message.
func Warningln(args ...interface{}) {
	logger.Warningln(
		au.BrightYellow(
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

// Errorf show a formatted log error message.
func Errorf(format string, args ...interface{}) {
	logger.Errorf(
		format,
		au.BrightRed(
			emoji.Sprintf("😡  "+format, args...),
		).BgBlack().Bold().Underline(),
	)
}

// Errorln show a formatted log error message.
func Errorln(args ...interface{}) {
	logger.Errorln(
		au.BrightRed(
			emoji.Sprintf("😡  " + fmt.Sprintf("%v", args...)),
		).BgBlack().Bold().Underline(),
	)
}
