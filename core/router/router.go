package router

import "github.com/LucasRodriguesOliveira/GoBones/internal/http"

type Router struct {
  routes map[string]Route
}

func New() *Router {
  return &Router{
    routes: map[string]Route{},
  }
}

func (r *Router) Register(
  path string,
  method string,
  handler http.HandlerFunc,
) {
  if r.routes == nil {
    r.routes = make(map[string]Route)
  }

  if r.routes[path] == nil {
    r.routes[path] = make(Route)
  }

  r.routes[path][method] = handler
}

func (r *Router) GetRoutes() map[string]Route {
  return r.routes
}
