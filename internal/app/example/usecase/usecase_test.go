package example

import (
	"fmt"
	"testing"

	example "app-module/internal/app/example/domain"
	mock "app-module/internal/app/example/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetExampleData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		input  int
		want   int
		result *example.Instance
		err    error
		expect bool
	}{
		{
			input: 1,
			want:  1,
			result: &example.Instance{
				ID:   1,
				Test: "1",
			},
			err:    nil,
			expect: true,
		},
		{
			input:  2,
			want:   2,
			result: nil,
			err:    fmt.Errorf("some error"),
			expect: false,
		},
	}

	for i := range tests {
		var (
			err    error
			result bool
			m      = mock.NewMockRepository(ctrl)
		)

		m.
			EXPECT().
			GetOne(tests[i].want).
			Return(tests[i].result, tests[i].err)

		_, err = New(m).GetExampleData(tests[i].input)
		if err == nil {
			result = true
		}

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}
