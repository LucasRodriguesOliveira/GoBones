package http

import (
	"fmt"
	"net/http"
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
    fmt.Printf("Something went wrong: %v", err);
  }
}
