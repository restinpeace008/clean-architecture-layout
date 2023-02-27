package app

import (
	example "app-module/internal/app/example"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() error {
	echo := echo.New()
	echo.HideBanner = true
	echo.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))
	api := echo.Group("/api")
	if err := example.Run(api); err != nil {
		return err
	}
	for _, kek := range echo.Router().Routes() {
		fmt.Println(kek.Path)
	}

	echo.Start("localhost:2222")
	return nil
}
