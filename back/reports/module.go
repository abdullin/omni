package reports

import (
	"net/http"

	"bitbucket.org/abdullin/proto/back/api"
	"bitbucket.org/abdullin/proto/back/module"
)

type Module struct{}

func (m *Module) Register(r module.Registrar) {
	r.HandleHttp("GET", "/inbox", m.listItems)
	r.HandleHttp("POST", "/inbox", m.addItem)
}

func (m *Module) listItems(r *api.Request) api.Response {
	return api.NewError("Not implemented", http.StatusOK)
}

func (m *Module) addItem(r *api.Request) api.Response {
	return api.NewError("Not implemented", http.StatusOK)
}

func NewModule() module.Module {
	return &Module{}
}

var Spec = &module.Spec{
	Name:    "inbox",
	Schema:  schema,
	Factory: NewModule,
}
