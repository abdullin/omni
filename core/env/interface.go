package module

import (
	"net/http"

	"github.com/abdullin/omni/core"
	"github.com/abdullin/omni/core/api"
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
	Name     string
	Factory  Factory
	UseCases []UseCaseFactory
}

type UseCaseFactory func() *UseCase

type EventHandler interface {
	Contracts() []string
	HandleEvent(e shared.Event) error
}

type EventHandlerMap map[string]EventHandler

type Publisher interface {
	Publish(e shared.Event) error
	MustPublish(e shared.Event)
}

type Request struct {
	Method  string
	Path    string
	Headers http.Header
	Body    interface{}
}

type Response struct {
	Status  int         `json:"status"`
	Headers http.Header `json:"headers"`
	Body    interface{} `json:"body"`
}

type UseCase struct {
	Name string

	Given []shared.Event
	When  *Request

	ThenEvents   []interface{}
	ThenResponse *Response
}
