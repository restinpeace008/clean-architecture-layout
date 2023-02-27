package errors

import (
	"fmt"
	"runtime"
)

type Error struct {
	text     string
	location string
}

func New(text string) error {
	var location *string
	_, file, line, ok := runtime.Caller(1)
	if ok {
		loc := fmt.Sprintf("%s:%d", file, line)
		location = &loc
	}
	return &Error{text: text, location: *location}
}

func (e *Error) Error() string {
	return e.text
}

func Wrap(e error, context string) error {
	if e == nil {
		return nil
	}

	var location *string
	_, file, line, ok := runtime.Caller(1)
	if ok {
		loc := fmt.Sprintf("%s:%d", file, line)
		location = &loc
	}

	return &Error{text: addWrap(e, context), location: *location}

}

func addWrap(e error, t string) string {
	return fmt.Sprintf("%s: %s", t, e.Error())
}
