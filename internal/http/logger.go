package http

import "log"

func Logger(req *Request, res *Response) error {
  log.Printf("[%s]: %s\n", req.Method, req.Path)

  return nil
}
