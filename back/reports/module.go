package reports

import (
	"net/http"

	"bitbucket.org/abdullin/proto/back/api"
	"bitbucket.org/abdullin/proto/back/events"
	"bitbucket.org/abdullin/proto/back/module"
)

type Module struct {
	pub module.Publisher
}

func (m *Module) Register(r module.Registrar) {
	r.HandleHttp("GET", "/reports/tx", m.listTransactions)

	var s = &store{}
	var d1 = &feedDenormalizer{s}

	r.HandleEvents("feed", d1)
}

func (m *Module) listTransactions(r *api.Request) api.Response {
	m.pub.MustPublish(&events.ProductCreated{})
	return api.NewError("Not implemented", http.StatusOK)
}

func NewModule(pub module.Publisher) module.Module {
	return &Module{pub}
}

var Spec = &module.Spec{
	Name:     "inbox",
	Schema:   schema,
	Factory:  NewModule,
	UseCases: cases,
}
