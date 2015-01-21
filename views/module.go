package views

import "github.com/abdullin/omni/core/env"

// basic module with views that don't fit anywhere else

type Module struct {
	pub env.Publisher
	d   *denormalizer
	s   *store
}

func NewModule(pub env.Publisher) *Module {
	store := newStore()
	denormalizer := newDenormalizer(store)
	return &Module{pub, denormalizer, store}
}

func (m *Module) Register(r env.Registrar) {
	r.HandleHttp("GET", "/views/inbox", m.getInbox)
	r.HandleEvents("views-denormalizer", m.d)
}

var Spec = &env.Spec{
	Name:     "views",
	UseCases: useCases,
}
