package http

type HandlerFunc func(*Request, *Response) error
