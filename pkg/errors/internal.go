package errors

type customError struct {
	text     string
	location string
	when     string
	wrapped  error
	codeHTTP int
}

func (e customError) Error() string {
	return e.text
}

func (e customError) getLocation() string {
	return e.location
}

func (e customError) getTime() string {
	return e.when
}

func (e customError) getWrapped() error {
	return e.wrapped
}

func (e customError) getCodeHTTP() int {
	return e.codeHTTP
}

func (e customError) setCodeHTTP(code int) error {
	e.codeHTTP = code
	return e
}

// WIP for What?
func (e customError) GetCode() int {
	return e.codeHTTP
}
