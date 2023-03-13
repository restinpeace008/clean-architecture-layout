package example

import (
	example "app-module/internal/app/example/domain"
	mock "app-module/internal/app/example/mock"
	"app-module/pkg/errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockUsecase(ctrl)

	tests := []example.TestCase{
		{
			Input: `{"some-id": 1}`,
			Want: map[string]example.DependencyMock{
				"GetExampleData": {
					Args: 1,
					Result: []any{
						&example.Instance{
							ID:   1,
							Test: "1",
						},
						nil,
					},
				},
			},
			Err: nil,
		},
		{
			Input: `{"some-id": 2}`,
			Want: map[string]example.DependencyMock{
				"GetExampleData": {
					Args: 2,
					Result: []any{
						nil,
						fmt.Errorf("some error"),
					},
				},
			},
			Err: fmt.Errorf("some error"),
		},
		{
			Input: `{"some-id": 0}`,
			Err:   fmt.Errorf("SomeID cannot be nil"),
		},
		{
			Input: `{"some-id": "0"}`,
			Err: fmt.Errorf("code=400, message=Unmarshal type error: expected=int, " +
				"got=string, field=some-id, offset=15, internal=json: cannot unmarshal " +
				"string into Go struct field Request.some-id of type int"),
		},
	}

	for i := range tests {
		if ucMock, ok := tests[i].Want["GetExampleData"]; ok {
			m.
				EXPECT().
				GetExampleData(ucMock.Args).
				Return(ucMock.Result...)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/test", strings.NewReader(tests[i].Input.(string)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		ctx := echo.New().NewContext(req, httptest.NewRecorder())
		d := &delivery{
			api: nil,
			uc:  m,
			l:   nil,
		}

		err := d.test(ctx)

		assert.True(t, errors.Is(err, tests[i].Err), "TestCase # %d", i+1)
	}
}
