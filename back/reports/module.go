package reports

import (
	"net/http"

	"bitbucket.org/abdullin/proto/back/api"
	"bitbucket.org/abdullin/proto/back/module"
)

type Module struct{}

func (m *Module) Register(r module.Registrar) {
	r.HandleHttp("GET", "/reports/tx", m.listTransactions)

	var s = &store{}
	var d1 = &feedDenormalizer{s}

	r.HandleEvents("feed", d1)
}

func (m *Module) listTransactions(r *api.Request) api.Response {
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
