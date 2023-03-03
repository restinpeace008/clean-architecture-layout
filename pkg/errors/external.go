package errors

import (
	"fmt"
	"runtime"
)

func New(text string, codeHTTP ...*int) error {
	var location string
	_, file, line, ok := runtime.Caller(1)
	if ok {
		location = fmt.Sprintf("%s:%d", file, line)
	}

	var code *int

	if len(codeHTTP) != 0 {
		code = codeHTTP[0]
	}

	return &customError{text: text, location: location, when: now(), codeHTTP: code}
}

func Wrap(e error, context string, codeHTTP ...*int) error {
	if e == nil {
		return nil
	}

	var location string

	_, file, line, ok := runtime.Caller(1)
	if ok {
		location = fmt.Sprintf("%s:%d", file, line)
	}

	var code *int

	if len(codeHTTP) != 0 {
		code = codeHTTP[0]
	}

	return &customError{text: addWrap(e, context), location: location, when: now(), wrapped: e, codeHTTP: code}
}

func Unwrap(err error) error {
	if e, ok := err.(interface{ getWrapped() error }); ok {
		return e.getWrapped()
	}
	return nil
}

func Location(err error) string {
	if e, ok := err.(interface{ getLocation() string }); ok {
		return e.getLocation()
	}
	return ""
}

func When(err error) string {
	if e, ok := err.(interface{ getTime() string }); ok {
		return e.getTime()
	}
	return ""
}

func Cause(err error) error {
	if err == nil {
		return nil
	}

	var wrapped error

	e, ok := err.(interface{ getWrapped() error })
	if ok {
		wrapped = e.getWrapped()
		if wrapped != nil {
			return Cause(wrapped)
		}
	}

	return err
}

func Is(err error, target error) bool {
	return Cause(err) == target
}

func CauseLocation(err error) string {
	if err == nil {
		return ""
	}

	if e, ok := err.(interface {
		getWrapped() error
		getLocation() string
	}); ok {
		if e.getLocation() == "" {
			return ""
		}

		if wrapped := e.getWrapped(); wrapped != nil {
			if er, ok := wrapped.(interface{ getLocation() string }); ok && er.getLocation() != "" {
				return CauseLocation(wrapped)
			}

			return e.getLocation()
		}

		return e.getLocation()
	}

	return ""
}

func GetCodeHTTP(err error) *int {
	if e, ok := err.(interface{ getCodeHTTP() *int }); ok {
		return e.getCodeHTTP()
	}
	return nil
}

func AddCodeHTTP(err error, code int) error {
	if e, ok := err.(interface{ setCodeHTTP(code *int) error }); ok {
		return e.setCodeHTTP(&code)
	}
	return err
}
