package logging

import (
	"github.com/lpmatos/loli/internal/helpers"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func setLogLevel(logLevel string) {
	if helpers.IsEmpty(logLevel) {
		logrus.SetLevel(defaultLogLevel)
	} else {
		if logLevel, err := logrus.ParseLevel(logLevel); err != nil {
			log.Fatalf("[LOLI] - Unable to parse logging level: %s\n", logLevel, err)
		} else {
			log.SetLevel(logLevel)
		}
	}
}

func setLogFormat(formatter string, colors bool) {
	switch formatter {
	case "json":
		logrus.SetFormatter(JSONFormatter(false))
	case "json-pretty":
		logrus.SetFormatter(JSONFormatter(true))
	case "text":
		logrus.SetFormatter(TextFormatter(colors))
	case "plain":
		plainFormatter := new(PlainFormatter)
		plainFormatter.TimestampFormat = timestampFormat
		plainFormatter.LevelDesc = levels
		logrus.SetFormatter(plainFormatter)
	default:
		logrus.SetFormatter(TextFormatter(colors))
	}
}
