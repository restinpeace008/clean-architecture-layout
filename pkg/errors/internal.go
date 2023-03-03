package errors

type customError struct {
	text     string
	location string
	when     string
	wrapped  error
}

func (e *customError) Error() string {
	return e.text
}

func (e *customError) getLocation() string {
	return e.location
}

func (e *customError) getTime() string {
	return e.when
}

func (e *customError) getWrapped() error {
	return e.wrapped
}
