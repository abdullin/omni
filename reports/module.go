package reports

import "github.com/abdullin/omni/module"

type Module struct {
	pub module.Publisher
}

func (m *Module) Register(r module.Registrar) {
}

func NewModule(pub module.Publisher) module.Module {
	return &Module{pub}
}

var Spec = &module.Spec{
	Name:     "reports",
	Factory:  NewModule,
	UseCases: cases,
}
