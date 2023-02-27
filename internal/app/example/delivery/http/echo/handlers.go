package example

import (
	example "app-module/internal/app/example/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (d Delivery) test(ctx echo.Context) error {
	var (
		data     *example.Instance
		err      error
		response example.Response
	)

	defer func() {
		if err != nil {
			response.Error = err.Error()
			response.Data = nil
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}
		response.Data = data
		ctx.JSON(http.StatusOK, response)
	}()

	if data, err = d.uc.GetExampleData(0); err != nil {
		return err
	}
	return nil
}
