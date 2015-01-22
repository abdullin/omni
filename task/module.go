package task

import "github.com/abdullin/omni/core/env"

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
	r.HandleHttp("POST", "/task", m.postTask)
	//r.HandleEvents("views-denormalizer", m.d)
	//r.ResetData("store", m.s.reset)
}

var Spec = &env.Spec{
	Name:     "task",
	UseCases: useCases,
}
