package module

import (
	"github.com/abdullin/omni/api"
)

func NewContainer() *Container {
	return &Container{
		Routes:   []*api.Route{},
		Handlers: EventHandlerMap{},
	}
}

type Container struct {
	Routes   []*api.Route
	Handlers EventHandlerMap
}

func (r *Container) HandleHttp(
	method string,
	path string,
	handler api.Handler) {
	r.Routes = append(r.Routes, api.NewRoute(method, path, handler))
}

func (r *Container) HandleEvents(
	name string,
	handler EventHandler,
) {
	r.Handlers[name] = handler
}
