package http

import "log"

func Logger(next HandlerFunc) HandlerFunc {
  return func(req *Request) *Response {
    log.Printf("[%s]: %s\n", req.Method, req.Path)
    return next(req)
  }
}
