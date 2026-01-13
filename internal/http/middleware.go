package http

type Middleware func(HandlerFunc) HandlerFunc
