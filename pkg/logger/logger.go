package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewFormatted() *logrus.Logger {
	return &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			// ForceFormatting: true,
		},
	}
}
