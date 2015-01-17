package setup

import (
	"bitbucket.org/abdullin/proto/back/api"

	"bitbucket.org/abdullin/proto/back/bus"

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

func (c *Context) WireHttp(router *mux.Router) {
	for _, x := range c.Items {
		for _, route := range x.Routes {
			api.Handle(router, route)
		}
	}
}

func (c *Context) WireHandlers(bus bus.Bus) {
	for _, x := range c.Items {
		for n, h := range x.Handlers {
			bus.AddEventHandler(n, h)
		}
	}
}

type Context struct {
	Items []*registrar
}

type registrar struct {
	Module   module.Module
	Routes   []*api.Route
	Handlers map[string]module.EventHandler
}

func newRegistrar() *registrar {
	return &registrar{
		Handlers: make(map[string]module.EventHandler),
	}
}

func (r *registrar) HandleHttp(
	method string,
	path string,
	handler api.Handler) {
	r.Routes = append(r.Routes, api.NewRoute(method, path, handler))
}

func (r *registrar) HandleEvents(
	name string,
	handler module.EventHandler,
) {
	r.Handlers[name] = handler
}
