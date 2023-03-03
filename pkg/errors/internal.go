package errors

type err struct {
	text     string
	location string
	when     string
	wrapped  error
	codeHTTP *int
}

func (e *err) Error() string {
	return e.text
}

func (e *err) getLocation() string {
	return e.location
}

func (e *err) getTime() string {
	return e.when
}

func (e *err) getWrapped() error {
	return e.wrapped
}

func (e *err) getCodeHTTP() *int {
	return e.codeHTTP
}

func (e *err) setCodeHTTP(code *int) error {
	e.codeHTTP = code
	return e
}
