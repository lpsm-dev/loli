package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	baseLogger := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}

	baseLogger.Out = os.Stdout
	baseLogger.Level = logrus.DebugLevel

	// Log as JSON instead of the default ASCII formatter.
	baseLogger.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	baseLogger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	baseLogger.SetLevel(logrus.WarnLevel)

	otherLogger := logrus.New()
	otherLogger.Level = logrus.DebugLevel
	otherLogger.SetReportCaller(true)

	baseLogger.Info("Base")
	otherLogger.Info("Other")

	otherLogger.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	otherLogger.Trace("Something very low level.")
	otherLogger.Debug("Useful debugging information.")
	otherLogger.Info("Something noteworthy happened!")
	otherLogger.Warn("You should probably take a look at this.")
	otherLogger.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	otherLogger.Fatal("Bye.")
	// Calls panic() after logging
	// otherLogger.Panic("I'm bailing.")
}
