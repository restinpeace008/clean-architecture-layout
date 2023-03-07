package app

import (
	example "app-module/internal/app/example"
	"app-module/pkg/errors"
	"app-module/pkg/logger"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Kinda boilerplate/global entrypoiint
func Start() error {
	// Create Echo instance
	echo := echo.New()
	// Disable framework banner
	echo.HideBanner = true

	// Create bare logger instance
	log := logger.Bare()

	// Attaching logging middleware globally
	echo.Use(logger.Middleware(log))

	// Launch entities constructs.
	if err := RunAll(echo, log); err != nil {
		return errors.Wrap(err, "App start")
	}

	// Start the server
	return echo.Start(fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")))
}

// The place where all of entities start
func RunAll(e *echo.Echo, l *logrus.Logger) error {
	// Create main `api` group which will contain some subgroups depends on entites count
	api := e.Group("/api")

	// Here for example provided 2 entities: `example` and `anotherExample`
	// In `example` case we don't provide any logger to this, because there is a feature: you can use any logging setup in each entity.
	if err := example.Run(api); err != nil {
		return err
	}

	// But in this case we gonna use already provided logger.
	// Honestly, I don't have any clue what's the point of it, but it wasn't hard to prepare this)'
	// Maybe it will be usefull for somebody.
	// Anyway, you (yeah, yeah - YOU) can edit this as you want.

	// if err := anotherExample.Run(api, l); err != nil {
	// 	return err
	// }

	return nil
}
