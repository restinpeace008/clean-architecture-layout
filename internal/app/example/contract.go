package example

import (
	d "app-module/internal/app/example/delivery/http/echo"
	r "app-module/internal/app/example/repository/some-source"
	uc "app-module/internal/app/example/usecase"
	"app-module/pkg/logger"
	"app-module/pkg/source"

	"github.com/labstack/echo/v4"
)

func Run(api *echo.Group) error {
	l := logger.NewFormatted()
	repository := r.New(
		l,
		source.New(),
	)
	usecase := uc.New(repository)
	delivery := d.New(api.Group("/example"), usecase, l)
	delivery.Expose()
	return nil
}
