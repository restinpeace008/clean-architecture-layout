package example

import (
	d "app-module/internal/app/example/delivery/http/echo"
	r "app-module/internal/app/example/repository/some-source"
	uc "app-module/internal/app/example/usecase"
	"app-module/pkg/logger"
	"app-module/pkg/source"

	"github.com/labstack/echo/v4"
)

// Run `example` construct
func Run(api *echo.Group) error {
	// Creating formatted logger
	l := logger.Formatted()

	// Init repository layer
	repository := r.New(l, source.New())

	// Init usecase layer
	usecase := uc.New(repository)

	// Init delivery layer
	delivery := d.New(api.Group("/example"), usecase, l)

	// Expose delivery routes
	delivery.Expose()

	// Error in this case just for the perspective.
	return nil
}
