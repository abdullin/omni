package setup

import (
	"bitbucket.org/abdullin/proto/back/api"
	"bitbucket.org/abdullin/proto/back/module"
	"github.com/gorilla/mux"
)

func modules(ms []module.Module) *Context {
	var c = &Context{}
	for _, m := range ms {
		r := newRegistrar()
		m.Register(r)

		c.Items = append(c.Items, r)
	}
	return c
}

func (c *Context) Wire(router *mux.Router) {
	for _, x := range c.Items {
		for _, route := range x.Routes {
			api.Handle(router, route)
		}
	}
}

type Context struct {
	Items []*registrar
}

type registrar struct {
	Module module.Module
	Routes []*api.Route
}

func newRegistrar() *registrar {
	return &registrar{}
}

func (r *registrar) HandleHttp(
	method string,
	path string,
	handler api.Handler) {
	r.Routes = append(r.Routes, api.NewRoute(method, path, handler))
}
