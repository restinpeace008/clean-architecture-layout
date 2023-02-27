package example

import (
	example "app-module/internal/app/example/domain"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Delivery struct {
	api *echo.Group
	uc  example.Usecase
	l   *logrus.Logger
}

func New(e *echo.Group, uc example.Usecase, l *logrus.Logger) example.Delivery {
	return Delivery{api: e, uc: uc, l: l}
}

func (d Delivery) Expose() {
	d.api.GET("/test", d.test)
}
