package errors

import (
	"fmt"
	"runtime"
)

func New(text string) error {
	var location string
	_, file, line, ok := runtime.Caller(1)
	if ok {
		location = fmt.Sprintf("%s:%d", file, line)
	}

	return &Error{text: text, location: location, when: now()}
}

func Wrap(err error, context string) error {
	if err == nil {
		return nil
	}

	var location string
	_, file, line, ok := runtime.Caller(1)
	if ok {
		location = fmt.Sprintf("%s:%d", file, line)
	}

	return &Error{text: addWrap(err, context), location: location, when: now(), wrapped: err}

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
