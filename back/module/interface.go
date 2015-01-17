package module

import (
	"bitbucket.org/abdullin/proto/back/api"
	"bitbucket.org/abdullin/proto/back/shared"
)

type Registrar interface {
	HandleHttp(method string, path string, handler api.Handler)
	HandleEvents(name string, handler EventHandler)
}

type Module interface {
	Register(r Registrar)
}

type Factory func(pub Publisher) Module

type Spec struct {
	Name    string
	Schema  string
	Factory Factory
}

type EventHandler interface {
	Contracts() []string
	HandleEvent(e shared.Event) error
}

type EventHandlerMap map[string]EventHandler

type Publisher interface {
	Publish(e shared.Event) error
	MustPublish(e shared.Event)
}
