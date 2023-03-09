package testcases

import (
	example "app-module/internal/app/example/domain"
	"fmt"
)

var (
	GetExampleDataTCases = []TestCase{
		{
			Input: 1,
			Want: map[string]DependencyMock{
				"r.GetOne": {
					Args: 1,
					Result: []any{
						&example.Instance{
							ID:   1,
							Test: "1",
						},
						nil,
					},
				},
				"sad.CheckSomeData": {
					Args:   "1",
					Result: []any{nil},
				},
			},
			Result: &example.Instance{
				ID:   1,
				Test: "1",
			},
			Err: nil,
		},
		{
			Input: 2,
			Want: map[string]DependencyMock{
				"r.GetOne": {
					Args: 2,
					Result: []any{
						nil,
						fmt.Errorf("some error"),
					},
				},
			},
			Result: nil,
			Err:    fmt.Errorf("some error"),
		},
		{
			Input: 3,
			Want: map[string]DependencyMock{
				"r.GetOne": {
					Args: 3,
					Result: []any{
						&example.Instance{
							ID:   3,
							Test: "3",
						},
						nil,
					},
				},
				"sad.CheckSomeData": {
					Args: "3",
					Result: []any{
						fmt.Errorf("some error"),
					},
				},
			},
			Result: nil,
			Err:    fmt.Errorf("some error"),
		},
	}
)
