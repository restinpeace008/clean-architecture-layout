package anotherOne

import (
	anotherOne "app-module/internal/app/anotherOne/domain"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// delivery instance
type delivery struct {
	api *echo.Group
	uc  anotherOne.Usecase
	l   *logrus.Logger
}

// New instance's factory
func New(e *echo.Group, uc anotherOne.Usecase, l *logrus.Logger) anotherOne.Delivery {
	// Inject dependencies
	return &delivery{api: e, uc: uc, l: l}
}

// Expose implementation
func (d *delivery) Expose() {
	// make some demo route
	d.api.POST("/test", d.test)
}
