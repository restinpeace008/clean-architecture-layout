package utils

import (
	"app-module/pkg/errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetStringPointer(s string) *string {
	return &s
}

func processError(err error) *string {
	if err != nil {
		return GetStringPointer(err.Error())
	}
	return nil
}

func processCode(err error) (code int) {
	code = http.StatusOK
	if err != nil {
		code = errors.CodeHTTP(err)
		if code == 0 {
			code = http.StatusInternalServerError
		}
	}
	return
}

func ProcessResponse(ctx echo.Context, l *logrus.Logger, data any, err error) {
	select {
	case <-ctx.Request().Context().Done():
		l.Warnln("[Example] Delivery:ProcessResponse client has closed connection")
		return
	default:
		if err != nil {
			data = nil
		}
		ctx.JSON(processCode(err), Response{Data: data, Error: processError(err)})
	}

}
