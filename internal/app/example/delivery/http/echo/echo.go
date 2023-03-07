package example

import (
	example "app-module/internal/app/example/domain"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// delivery instance
type delivery struct {
	api *echo.Group
	uc  example.Usecase
	l   *logrus.Logger
}

// New instance's factory
func New(e *echo.Group, uc example.Usecase, l *logrus.Logger) example.Delivery {
	// Inject dependencies
	return &delivery{api: e, uc: uc, l: l}
}

// Expose implementation
func (d *delivery) Expose() {
	// make some demo route
	d.api.GET("/test", d.test)
}
