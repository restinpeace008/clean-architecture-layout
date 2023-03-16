package anotherOne

import (
	anotherOne "app-module/internal/app/anotherOne/domain"
	"app-module/pkg/errors"
	"app-module/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// test demo endpoint handler
func (d *delivery) test(ctx echo.Context) error {
	var (
		request  anotherOne.Request
		response *anotherOne.Response
		data     *anotherOne.Instance
		err      error
	)

	// This construction allows you to send universal answers to your clients
	defer func() {
		utils.ProcessResponse(ctx, d.l, response, err)
	}()

	// Unserialize user input
	// Note: we should wrap it by out pkg/errors, because it's library method (ctx.Bind)
	if err = errors.Wrap(ctx.Bind(&request), "test: unserialize error", http.StatusBadRequest); err != nil {
		return err
	}

	// Validate user input
	// But here, unlike `ctx.Bind`, we can inject http code inside `Validate` method (see inside).
	if err = request.Validate(); err != nil {
		return err
	}

	// We use pointer to `anotherOne.Response` for provide `null` in answer, this is why this line exists.
	response = &anotherOne.Response{SomeData: data}

	return nil
}
