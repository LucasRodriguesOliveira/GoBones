package http

import (
	"fmt"
	"net/http"
	"strings"
)

type Response struct {
	Status  int
	Headers map[string]string
	Body    []byte
}

func WriteResponse(w http.ResponseWriter, res *Response) {
	for k, v := range res.Headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(res.Status)
	if _, err := w.Write(res.Body); err != nil {
		fmt.Printf("Something went wrong: %v", err)
	}
}

func (r *Response) jsonHeader() {
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}
	r.Headers["Content-Type"] = "Application/json"
}

func (r *Response) jsonBody(body J) {
	r.jsonHeader()
	content := []string{"{"}

	for k, v := range body {
		switch v.(type) {
		case string:
			content = append(content, fmt.Sprintf("%q: %q", k, v))
		default:
			content = append(content, fmt.Sprintf("%q: %v", k, v))
		}
	}

	content = append(content, "}")

	r.Body = fmt.Appendln([]byte(strings.Join(content, ",")))
}

func (r *Response) Ok(body J) {
	r.Status = http.StatusOK
	r.jsonBody(body)
}

func (r *Response) Created(body J) {
	r.Status = http.StatusCreated
	r.jsonBody(body)
}
