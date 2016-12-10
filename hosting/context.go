package hosting

import (
	"github.com/abdullin/omni/core/api"

	"github.com/abdullin/omni/core/bus"

	"github.com/abdullin/omni/core/env"
	"github.com/gorilla/mux"
)

func modules(ms []env.Module) *Context {
	var c = &Context{}
	for _, m := range ms {
		r := env.NewContainer()
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
	Items []*env.Container
}
