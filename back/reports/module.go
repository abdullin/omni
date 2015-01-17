package reports

import (
	"net/http"

	"bitbucket.org/abdullin/proto/back/api"
	"bitbucket.org/abdullin/proto/back/events"
	"bitbucket.org/abdullin/proto/back/module"
	"bitbucket.org/abdullin/proto/back/shared"
)

type Module struct{}

func (m *Module) Register(r module.Registrar) {
	r.HandleHttp("GET", "/reports/transactions", m.listTransactions)

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

type store struct {
	Events []shared.Event
}

func (s *store) AddEvent(e shared.Event) {
	s.Events = append(s.Events, e)
}

type feedDenormalizer struct {
	store *store
}

func (f *feedDenormalizer) Contracts() []string {
	return []string{
		"ProductCreated",
	}
}
func (f *feedDenormalizer) HandleEvent(e shared.Event) error {
	switch t := e.(type) {
	case *events.ProductCreated:
		f.store.AddEvent(t)
	}
	return nil

}
