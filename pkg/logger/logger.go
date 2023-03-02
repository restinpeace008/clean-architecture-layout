package logger

import (
	"app-module/pkg/errors"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var Default *logrus.Logger

// Allows to use global package logger.
// if `isBare` == true then default logrus will be used
// Otherwise here will be used formatted logger. It's cannot be configured yet. (TODO)
// import `logger` package and use like this: `logger.Default.Infoln("tell me smthng") e.g..`
func UseGlobal(isBare bool) {
	if isBare {
		Default = Bare()
		return
	}
	Default = Formatted()
}

// Returns absolutely bare *logrus.Logger
func Bare() *logrus.Logger {
	return logrus.New()
}

// Returns a bit configured logger. (TODO: add configuring if you want)
func Formatted() *logrus.Logger {
	return &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			// ForceFormatting: true,
		},
	}
}

// Logging middleware for echo framework
// This works with custom errors package. See pkg/errors.
func Middleware(log *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err error

			start := time.Now().UTC()

			if err = next(c); err != nil {
				c.Error(err)
			}

			res := c.Response()

			fields := logrus.Fields{
				"status":  res.Status,
				"latency": time.Since(start).String(),
				"user_ip": c.RealIP(),
			}

			if err != nil {
				fields["when"] = errors.When(err)
				fields["location"] = errors.Location(err)
				fields["error"] = err.Error()
			}

			switch {
			case res.Status >= 500:
				log.WithFields(fields).Info("Server error")
			case res.Status >= 400:
				log.WithFields(fields).Info("Client error")
			case res.Status >= 300:
				log.WithFields(fields).Info("Redirection")
			default:
				log.WithFields(fields).Info("Success")
			}

			return nil
		}
	}
}
