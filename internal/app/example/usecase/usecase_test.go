package example

import (
	"testing"

	example "app-module/internal/app/example/domain"
	mock "app-module/internal/app/example/mock"
	"app-module/pkg/errors"

	"app-module/internal/app/example/usecase/testcases"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetExampleData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for i, tcase := range testcases.GetExampleDataTCases {
		r := mock.NewMockRepository(ctrl)
		s := mock.NewMockSomeService(ctrl)

		if repoMock, ok := tcase.Want["r.GetOne"]; ok {
			r.
				EXPECT().
				GetOne(repoMock.Args).
				Return(repoMock.Result...)
		}

		if serviceMock, ok := tcase.Want["s.CheckSomeData"]; ok {
			s.
				EXPECT().
				CheckSomeData(serviceMock.Args).
				Return(serviceMock.Result...)
		}

		id, ok := tcase.Input.(int)
		if !ok {
			assert.FailNow(t, "invalid input", "TestCase # %d", i+1)
		}

		var expectedResult *example.Instance

		if tcase.Result != nil {
			if expectedResult, ok = tcase.Result.(*example.Instance); !ok {
				assert.FailNow(t, "invalid expected result", "TestCase # %d", i+1)
			}
		}

		result, err := New(r, s).GetExampleData(id)

		assert.Equal(t, result, expectedResult, "TestCase # %d", i+1)
		assert.Equal(t, errors.Cause(err), tcase.Err, "TestCase # %d", i+1)
	}
}
