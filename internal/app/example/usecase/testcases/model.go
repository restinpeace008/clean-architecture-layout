package testcases

type (
	DependencyMock struct {
		Args   any
		Result []any // []any{result, error} error is always at the end
	}

	TestCase struct {
		Input  any                       // args for usecase func
		Want   map[string]DependencyMock // mocks for repository and delivery interfaces
		Result any                       // expected result (if exists) for usecase func
		Err    error                     // expected error for usecase func
	}
)
