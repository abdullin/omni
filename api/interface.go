package api

import "net/http"

type Handler func(r *Request) Response

type Route struct {
	Method  string
	Path    string
	Handler Handler
}

func NewRoute(method, path string, handler Handler) *Route {
	return &Route{method, path, handler}
}

type Response interface {
	Write(w http.ResponseWriter)
}
