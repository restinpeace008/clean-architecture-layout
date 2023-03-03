package errors

import (
	"fmt"
	"runtime"
)

func New(text string, codeHTTP ...interface{}) error {
	var location string
	_, file, line, ok := runtime.Caller(1)
	if ok {
		location = fmt.Sprintf("%s:%d", file, line)
	}

	var code *int
	if val, ok := codeHTTP[0].(*int); ok {
		code = val
	}
	return &err{text: text, location: location, when: now(), codeHTTP: code}
}

func Wrap(e error, context string, codeHTTP ...interface{}) error {
	if e == nil {
		return nil
	}

	var location string
	_, file, line, ok := runtime.Caller(1)
	if ok {
		location = fmt.Sprintf("%s:%d", file, line)
	}

	var code *int
	if val, ok := codeHTTP[0].(*int); ok {
		code = val
	}

	return &err{text: addWrap(e, context), location: location, when: now(), wrapped: e, codeHTTP: code}

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
