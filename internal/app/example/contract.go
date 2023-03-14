package example

import (
	d "app-module/internal/app/example/delivery/http/echo"
	r "app-module/internal/app/example/infrastructure/repository/postgres"
	s "app-module/internal/app/example/infrastructure/service/someapi"
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

	// Init service layer
	someService := s.New("")

	// Init usecase layer
	usecase := uc.New(repository, someService)

	// Init delivery layer
	delivery := d.New(api.Group("/example"), usecase, l)

	// Expose delivery routes
	delivery.Expose()

	// Error in this case just for the perspective.
	return nil
}
