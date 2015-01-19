package reports

import (
	"bitbucket.org/abdullin/proto/back/api"
	"bitbucket.org/abdullin/proto/back/events"
	"bitbucket.org/abdullin/proto/back/module"
)

type Module struct {
	pub module.Publisher
}

func (m *Module) Register(r module.Registrar) {
	r.HandleHttp("GET", "/reports/groups", m.listGroups)

	var s = &store{}
	var d1 = &feedDenormalizer{s}

	r.HandleEvents("feed", d1)
}

func (m *Module) listGroups(r *api.Request) api.Response {
	// wild response
	m.pub.MustPublish(&events.ProductCreated{})

	return api.NewJSON(&GroupReportPageDto{
		Items: []*GroupReportItem{
			&GroupReportItem{
				Quantity: 3,
				Name:     "hardcoded",
			},
		},
	})
}

func NewModule(pub module.Publisher) module.Module {
	return &Module{pub}
}

var Spec = &module.Spec{
	Name:     "inbox",
	Factory:  NewModule,
	UseCases: cases,
}
