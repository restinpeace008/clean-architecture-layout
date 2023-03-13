package errors

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// internal.Error()
func TestError(t *testing.T) {
	tests := []struct {
		input  string
		want   string
		expect bool
	}{
		{
			input:  "foo",
			want:   "foo",
			expect: true,
		},
		{
			input:  "foo",
			want:   "bar",
			expect: false,
		},
	}

	for i := range tests {
		result := New(tests[i].input).Error() == tests[i].want

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

// internal.getLocation()
func TestGetLocation(t *testing.T) {
	first := New("first")
	second := New("second")

	f := first.(interface{ getLocation() string })
	s := second.(interface{ getLocation() string })

	tests := []struct {
		input  error
		want   string
		expect bool
	}{
		{
			input:  first,
			want:   f.getLocation(),
			expect: true,
		},
		{
			input:  first,
			want:   s.getLocation(),
			expect: false,
		},
		{
			input:  second,
			want:   "",
			expect: false,
		},
	}

	for i := range tests {
		input, ok := tests[i].input.(interface{ getLocation() string })
		if !ok {
			assert.FailNow(t, "invalid input", "TestCase # %d", i+1)
		}

		result := input.getLocation() == tests[i].want

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

// internal.getTime()
func TestGetTime(t *testing.T) {
	first := New("first")
	second := New("second")

	f := first.(interface{ getTime() string })
	s := second.(interface{ getTime() string })

	tests := []struct {
		input  error
		want   string
		expect bool
	}{
		{
			input:  first,
			want:   f.getTime(),
			expect: true,
		},
		{
			input:  first,
			want:   s.getTime(),
			expect: false,
		},
		{
			input:  second,
			want:   "",
			expect: false,
		},
	}

	for i := range tests {
		input, ok := tests[i].input.(interface{ getTime() string })
		if !ok {
			assert.FailNow(t, "invalid input", "TestCase # %d", i+1)
		}

		result := input.getTime() == tests[i].want

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

// internal.getWrapped()
func TestGetWrapped(t *testing.T) {
	first := New("first")
	second := New("second")

	wrappedFirst := Wrap(first, "1")
	wrappedSecond := Wrap(second, "1")

	tests := []struct {
		input  error
		want   error
		expect bool
	}{
		{
			input:  wrappedFirst,
			want:   first,
			expect: true,
		},
		{
			input:  second,
			want:   nil,
			expect: true,
		},
		{
			input:  wrappedFirst,
			want:   second,
			expect: false,
		},
		{
			input:  wrappedSecond,
			want:   nil,
			expect: false,
		},
	}

	for i := range tests {
		input, ok := tests[i].input.(interface{ getWrapped() error })
		if !ok {
			assert.FailNow(t, "invalid input", "TestCase # %d", i+1)
		}

		result := input.getWrapped() == tests[i].want

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

// internal.getCodeHTTP()
func TestGetCodeHTTP(t *testing.T) {
	first := New("first", http.StatusBadRequest)
	second := New("second", http.StatusForbidden)

	tests := []struct {
		input  error
		want   int
		expect bool
	}{
		{
			input:  first,
			want:   http.StatusBadRequest,
			expect: true,
		},
		{
			input:  second,
			want:   http.StatusAccepted,
			expect: false,
		},
	}

	for i := range tests {
		input, ok := tests[i].input.(interface{ getCodeHTTP() int })
		if !ok {
			assert.FailNow(t, "invalid input", "TestCase # %d", i+1)
		}

		result := input.getCodeHTTP() == tests[i].want

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

// internal.setCodeHTTP()
func TestSetCodeHTTP(t *testing.T) {
	e := New("new error").(interface{ setCodeHTTP(int) error })
	err := e.setCodeHTTP(http.StatusBadRequest)

	tests := []struct {
		input  error
		want   int
		expect bool
	}{
		{
			input:  err,
			want:   http.StatusBadRequest,
			expect: true,
		},
		{
			input:  err,
			want:   http.StatusOK,
			expect: false,
		},
	}

	for i := range tests {
		input, ok := tests[i].input.(interface{ getCodeHTTP() int })
		if !ok {
			assert.FailNow(t, "invalid input", "TestCase # %d", i+1)
		}

		result := input.getCodeHTTP() == tests[i].want

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

// external.New()
func TestNew(t *testing.T) {
	errMessage := "new error"

	tests := []struct {
		input  string
		code   int
		expect bool
	}{
		{
			input:  errMessage,
			expect: true,
		},
		{
			input:  errMessage,
			code:   400,
			expect: true,
		},
		{
			input:  "",
			expect: true,
		},
	}

	for i := range tests {
		err := New(tests[i].input, tests[i].code)
		if err == nil {
			assert.FailNow(t, "broken func", "TestCase # %d", i+1)
		}

		if _, ok := err.(customError); !ok {
			assert.FailNow(t, "broken result", "TestCase # %d", i+1)
		}

		assert.Equal(t, true, true)
	}
}

// external.Wrap()
func TestWrap(t *testing.T) {
	testMsg := "new err"
	testErr := New(testMsg)

	tests := []struct {
		context string
		wrapped error
		code    int
		want    string
		expect  bool
	}{
		{
			context: "1",
			wrapped: testErr,
			want:    addWrap(testErr, "1"),
			expect:  true,
		},
		{
			context: "1",
			wrapped: testErr,
			code:    400,
			want:    addWrap(testErr, "1"),
			expect:  true,
		},
		{
			context: "1",
			wrapped: nil,
			want:    "",
			expect:  true,
		},
		{
			context: "1",
			wrapped: testErr,
			want:    addWrap(testErr, "2"),
			expect:  false,
		},
		{
			context: "1",
			wrapped: nil,
			want:    "1",
			expect:  false,
		},
	}

	for i := range tests {
		var msg string

		err := Wrap(tests[i].wrapped, tests[i].context, tests[i].code)
		if err != nil {
			if _, ok := err.(customError); !ok {
				assert.FailNow(t, "broken result", "TestCase # %d", i+1)
			}

			msg = err.Error()
		}

		result := msg == tests[i].want

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

// external.Unwrap()
func TestUnwrap(t *testing.T) {
	testMsg := "new err"
	testErr := New(testMsg)

	tests := []struct {
		input  error
		want   error
		expect bool
	}{
		{
			input:  Wrap(testErr, "1"),
			want:   testErr,
			expect: true,
		},
		{
			input:  testErr,
			want:   nil,
			expect: true,
		},
		{
			input:  Wrap(Wrap(testErr, "1"), "2"),
			want:   testErr,
			expect: false,
		},
	}

	for i := range tests {
		result := tests[i].want == Unwrap(tests[i].input)

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

// external.Location()
func TestLocation(t *testing.T) {
	first := New("first")
	second := Wrap(first, "second")

	f := first.(interface{ getLocation() string })
	s := second.(interface{ getLocation() string })

	tests := []struct {
		input  error
		want   string
		expect bool
	}{
		{
			input:  second,
			want:   s.getLocation(),
			expect: true,
		},
		{
			input:  nil,
			want:   "",
			expect: true,
		},
		{
			input:  second,
			want:   f.getLocation(),
			expect: false,
		},
	}

	for i := range tests {
		result := Location(tests[i].input) == tests[i].want

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

// external.When()
func TestWhen(t *testing.T) {
	first := New("first")
	second := Wrap(first, "second")

	f := first.(interface{ getTime() string })
	s := second.(interface{ getTime() string })

	tests := []struct {
		input  error
		want   string
		expect bool
	}{
		{
			input:  second,
			want:   s.getTime(),
			expect: true,
		},
		{
			input:  nil,
			want:   "",
			expect: true,
		},
		{
			input:  second,
			want:   f.getTime(),
			expect: false,
		},
	}

	for i := range tests {
		result := When(tests[i].input) == tests[i].want

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

// external.Cause()
func TestCause(t *testing.T) {
	customErr := New("custom error")
	defaultErr := fmt.Errorf("default error")

	tests := []struct {
		input  error
		want   error
		expect bool
	}{
		{
			input:  customErr,
			want:   customErr,
			expect: true,
		},
		{
			input:  Wrap(Wrap(customErr, "1"), "2"),
			want:   customErr,
			expect: true,
		},
		{
			input:  defaultErr,
			want:   defaultErr,
			expect: true,
		},
		{
			input:  Wrap(Wrap(defaultErr, "1"), "2"),
			want:   defaultErr,
			expect: true,
		},
		{
			input:  nil,
			want:   nil,
			expect: true,
		},
		{
			input:  Wrap(Wrap(defaultErr, "1"), "2"),
			want:   Wrap(defaultErr, "1"),
			expect: false,
		},
	}

	for i := range tests {
		result := Cause(tests[i].input) == tests[i].want

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

// external.Is()
func TestIs(t *testing.T) {
	customErr := New("custom error")
	defaultErr := fmt.Errorf("default error")

	tests := []struct {
		input  error
		want   error
		expect bool
	}{
		{
			input:  customErr,
			want:   customErr,
			expect: true,
		},
		{
			input:  Wrap(Wrap(customErr, "1"), "2"),
			want:   customErr,
			expect: true,
		},
		{
			input:  defaultErr,
			want:   defaultErr,
			expect: true,
		},
		{
			input:  New("custom error"),
			want:   customErr,
			expect: true,
		},
		{
			input:  Wrap(Wrap(defaultErr, "1"), "2"),
			want:   defaultErr,
			expect: true,
		},
		{
			input:  Wrap(Wrap(defaultErr, "1"), "2"),
			want:   Wrap(defaultErr, "1"),
			expect: false,
		},
		{
			input:  Wrap(Wrap(New("custom error"), "1"), "2"),
			want:   Wrap(customErr, "1"),
			expect: false,
		},
	}

	for i := range tests {
		result := Is(tests[i].input, tests[i].want)

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

// external.CauseLocation()
func TestCauseLocation(t *testing.T) {
	customCause := New("custom error")
	defaultCause := fmt.Errorf("default error")
	defaultWrapped := Wrap(defaultCause, "1")

	cc := customCause.(interface{ getLocation() string })
	dw := defaultWrapped.(interface{ getLocation() string })

	tests := []struct {
		input  error
		want   string
		expect bool
	}{
		{
			input:  Wrap(Wrap(customCause, "1"), "2"),
			want:   cc.getLocation(),
			expect: true,
		},
		{
			input:  defaultCause,
			want:   "",
			expect: true,
		},
		{
			input:  customCause,
			want:   cc.getLocation(),
			expect: true,
		},
		{
			input:  Wrap(defaultWrapped, "1"),
			want:   dw.getLocation(),
			expect: true,
		},
		{
			input:  nil,
			want:   "",
			expect: true,
		},
		{
			input:  fmt.Errorf("1 -> %w", defaultWrapped),
			want:   dw.getLocation(),
			expect: false,
		},
	}

	for i := range tests {
		result := CauseLocation(tests[i].input) == tests[i].want

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

// internal.CodeHTTP()
func TestCodeHTTP(t *testing.T) {
	first := AddCodeHTTP(New("first"), http.StatusBadRequest)
	second := AddCodeHTTP(New("second"), http.StatusForbidden)

	tests := []struct {
		input  error
		want   int
		expect bool
	}{
		{
			input:  first,
			want:   http.StatusBadRequest,
			expect: true,
		},
		{
			input:  second,
			want:   http.StatusAccepted,
			expect: false,
		},
	}

	for i := range tests {
		result := CodeHTTP(tests[i].input) == tests[i].want

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}

// internal.AddCodeHTTP()
func TestAddCodeHTTP(t *testing.T) {
	err := New("new error")

	tests := []struct {
		input  error
		code   int
		want   int
		expect bool
	}{
		{
			input:  err,
			code:   http.StatusBadRequest,
			want:   http.StatusBadRequest,
			expect: true,
		},
		{
			input:  err,
			code:   http.StatusBadRequest,
			want:   http.StatusOK,
			expect: false,
		},
	}

	for i := range tests {
		e := AddCodeHTTP(err, http.StatusBadRequest)

		result := CodeHTTP(e) == tests[i].want

		assert.Equal(t, result, tests[i].expect, "TestCase # %d", i+1)
	}
}
