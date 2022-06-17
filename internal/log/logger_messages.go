package log

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
	au "github.com/logrusorgru/aurora/v3"
)

// Debug function to print a pretty formatted log debug message
func Debug(args ...interface{}) {
	logger.Debug(
		au.BrightCyan(
			emoji.Sprintf("🆗  " + fmt.Sprintf("%v", args...)),
		).Bold().Underline(),
	)
}

// Debugf function to print a pretty formatted log debug message
func Debugf(format string, args ...interface{}) {
	logger.Debugf(
		"🆗 "+au.BrightCyan(format).
			Bold().
			Underline().
			String(), args...,
	)
}

// Debugln function to print a pretty formatted log debug message
func Debugln(args ...interface{}) {
	logger.Debugln(
		au.BrightCyan(
			emoji.Sprintf("🆗  " + fmt.Sprintf("%v\n", args...)),
		).Bold().Underline(),
	)
}

// Info function to print a pretty formatted log info message
func Info(args ...interface{}) {
	logger.Info(
		au.Green(
			emoji.Sprintf("✅  " + fmt.Sprintf("%v", args...)),
		).Bold().Underline(),
	)
}

// Infof function to print a pretty formatted log info message
func Infof(format string, args ...interface{}) {
	logger.Infof(
		"✅ "+au.Green(format).
			Bold().
			Underline().
			String(), args...,
	)
}

// Infoln function to print a pretty formatted log info message
func Infoln(args ...interface{}) {
	logger.Infoln(
		au.Green(
			emoji.Sprintf("✅  " + fmt.Sprintf("%v\n", args...)),
		).Bold().Underline(),
	)
}

// Warn function to print a pretty formatted log warn message
func Warn(args ...interface{}) {
	logger.Warn(
		au.BrightYellow(
			emoji.Sprintf("😲  " + fmt.Sprintf("%v", args...)),
		).Bold().Underline(),
	)
}

// Warnf function to print a pretty formatted log warn message
func Warnf(format string, args ...interface{}) {
	logger.Warnf(
		"😲 "+au.BrightYellow(format).
			Bold().
			Underline().
			String(), args...,
	)
}

// Warnln function to print a pretty formatted log warn message
func Warnln(args ...interface{}) {
	logger.Warnln(
		au.BrightYellow(
			emoji.Sprintf("😲  " + fmt.Sprintf("%v\n", args...)),
		).Bold().Underline(),
	)
}

// Error function to print a pretty formatted log error message
func Error(args ...interface{}) {
	logger.Error(
		au.BrightRed(
			emoji.Sprintf("😡  " + fmt.Sprintf("%v", args...)),
		))
}

// Errorf function to print a pretty formatted log error message
func Errorf(format string, args ...interface{}) {
	logger.Errorf(
		"😡 "+au.BrightRed(format).
			Bold().
			Underline().
			String(), args...,
	)
}

// Errorln function to print a pretty formatted log error message
func Errorln(args ...interface{}) {
	logger.Errorln(
		au.BrightRed(
			emoji.Sprintf("😡  " + fmt.Sprintf("%v\n", args...)),
		).Bold().Underline(),
	)
}

// Fatal function to print a pretty formatted log fatal message
func Fatal(args ...interface{}) {
	logger.Fatal(
		au.BrightRed(
			emoji.Sprintf("🤬  " + fmt.Sprintf("%v", args...)),
		))
}

// Fatalf function to print a pretty formatted log fatal message
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(
		"🤬 "+au.BrightRed(format).
			Bold().
			Underline().
			String(), args...,
	)
}

// Fatalln function to print a pretty formatted log fatal message
func Fatalln(args ...interface{}) {
	logger.Fatalln(
		au.BrightRed(
			emoji.Sprintf("🤬  " + fmt.Sprintf("%v\n", args...)),
		).Bold().Underline(),
	)
}
