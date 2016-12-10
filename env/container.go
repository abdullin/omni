package env

import (
	"github.com/abdullin/omni/api"
)

func NewContainer() *Container {
	return &Container{
		Routes:    []*api.Route{},
		Handlers:  EventHandlerMap{},
		DataReset: make(map[string]func()),
	}
}

type Container struct {
	Routes    []*api.Route
	Handlers  EventHandlerMap
	DataReset map[string]func()
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

func (r *Container) ResetData(name string, reset func()) {
	r.DataReset[name] = reset
}
