package example

import (
	d "app-module/internal/app/example/delivery/http/echo"
	sad "app-module/internal/app/example/delivery/http/someapi"
	r "app-module/internal/app/example/repository/postgres"
	uc "app-module/internal/app/example/usecase"
	"app-module/pkg/logger"
	"app-module/pkg/postgres"

	"github.com/labstack/echo/v4"
)

// Run `example` construct
func Run(api *echo.Group) error {
	// Creating formatted logger
	l := logger.Formatted()

	// Init repository layer
	repository := r.New(l, postgres.New())

	// Init delivery layer
	someApiDelivery := sad.New("")

	// Init usecase layer
	usecase := uc.New(repository, someApiDelivery)

	// Init delivery layer
	delivery := d.New(api.Group("/example"), usecase, l)

	// Expose delivery routes
	delivery.Expose()

	// Error in this case just for the perspective.
	return nil
}
