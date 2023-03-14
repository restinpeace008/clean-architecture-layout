package someapi

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// RoundTripFunc custom type for mock transport
type roundTripFunc func(req *http.Request) *http.Response

// RoundTrip interface implementation for http.RoundTripper.
func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func TestCheckSomeData(t *testing.T) {
	url := "https://example.com"

	tests := []struct {
		input  string
		param  string
		status int
		exists bool
	}{
		{
			input:  "some/path",
			param:  "some/path",
			status: 200,
			exists: true,
		},
		{
			input:  "some/path",
			param:  "some/path",
			status: 400,
			exists: false,
		},
		{
			input:  "some/path\\^%",
			param:  "some/path\\^%",
			exists: false,
		},
	}

	for i := range tests {
		// set custom default client
		http.DefaultClient = &http.Client{
			Transport: roundTripFunc(func(req *http.Request) *http.Response {
				assert.Equal(t, req.URL.String(), url+"/"+tests[i].param, "TestCase # %d", i+1)

				return &http.Response{
					StatusCode: tests[i].status,
					// Send response to be tested
					Body:   io.NopCloser(bytes.NewBufferString(`OK`)),
					Header: make(http.Header),
				}
			}),
		}

		err := New(url).CheckSomeData(tests[i].input)

		assert.Equal(t, tests[i].exists, err == nil, "TestCase # %d", i+1)
	}
}
