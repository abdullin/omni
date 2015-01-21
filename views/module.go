package views

import "github.com/abdullin/omni/core/env"

// basic module with views that don't fit anywhere else

type Module struct {
	pub env.Publisher
}

func NewModule(pub env.Publisher) *Module {
	return &Module{pub}
}

func (m *Module) Register(r env.Registrar) {
	r.HandleHttp("GET", "/views/inbox", m.getInbox)
}

var Spec = &env.Spec{
	Name:     "views",
	UseCases: useCases,
}
