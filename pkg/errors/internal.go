package errors

type Error struct {
	text     string
	location string
	when     string
	wrapped  error
}

func (e *Error) Error() string {
	return e.text
}

func (e *Error) getLocation() string {
	return e.location
}

func (e *Error) getTime() string {
	return e.when
}

func (e *Error) getWrapped() error {
	return e.wrapped
}
