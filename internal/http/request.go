package http

import (
	"io"
	"net/http"
)

type Request struct {
	Method  string
	Path    string
	Headers map[string][]string
	Body    []byte
}

func ParseRequest(r *http.Request) (*Request, error) {
	b, err := io.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	return &Request{
		Method:  r.Method,
		Path:    r.URL.Path,
		Headers: r.Header,
		Body:    b,
	}, nil
}
