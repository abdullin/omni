package reports

import "github.com/abdullin/omni/core/env"

type Module struct {
	pub env.Publisher
}

func (m *Module) Register(r env.Registrar) {
}

func NewModule(pub env.Publisher) env.Module {
	return &Module{pub}
}

var Spec = &env.Spec{
	Name:     "reports",
	Factory:  NewModule,
	UseCases: cases,
}
