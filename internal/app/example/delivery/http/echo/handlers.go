package example

import (
	example "app-module/internal/app/example/domain"
	"app-module/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// test demo endpoint handler
func (d Delivery) test(ctx echo.Context) error {
	var (
		data     *example.Instance
		err      error
		response example.Response
	)

	// This construction allows you to send universal answers to your clients
	defer func() {
		response.Error = utils.ProcessError(err)
		response.Data = data
		ctx.JSON(http.StatusInternalServerError, response) // TODO: think about status codes.
	}()

	// Here is depending place. Just call some `usecase` method.
	if data, err = d.uc.GetExampleData(0); err != nil {
		return err
	}

	return nil
}
