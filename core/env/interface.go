package env

import (
	"net/http"

	"github.com/abdullin/omni/core"
	"github.com/abdullin/omni/core/api"
)

type Registrar interface {
	HandleHttp(method string, path string, handler api.Handler)
	HandleEvents(name string, handler EventHandler)
	ResetData(name string, action func())
}

type Module interface {
	Register(r Registrar)
	//	Spec() *Spec
}

type Spec struct {
	Name string
	//Factory  Factory
	UseCases []UseCaseFactory
}

type UseCaseFactory func() *UseCase

type EventHandler interface {
	Contracts() []string
	HandleEvent(e core.Event) error
}

type EventHandlerMap map[string]EventHandler

type Publisher interface {
	Publish(e ...core.Event) error
	MustPublish(e ...core.Event)
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

	Given []core.Event
	When  *Request

	ThenEvents   []core.Event
	ThenResponse *Response

	Where Where
}

type Where interface {
	Map() map[interface{}]string
}
