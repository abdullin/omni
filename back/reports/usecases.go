package reports

import (
	"bitbucket.org/abdullin/proto/back/events"
	"bitbucket.org/abdullin/proto/back/module"
	"bitbucket.org/abdullin/proto/back/shared"
	"bitbucket.org/abdullin/proto/back/spec"
	"github.com/abdullin/seq"
)

var cases = []module.UseCaseFactory{
	First,
}

func First() *module.UseCase {
	p1 := events.NewProductCreated(id(), prod(), "pencil")
	p2 := events.NewProductCreated(id(), prod(), "notepad")

	l1 := events.NewLocationCreated(id(), loc(), "loc1")

	i1 := events.NewItemAdded(id(), p1.ProductId, l1.LocationId, 10)
	i2 := events.NewItemAdded(id(), p2.ProductId, l1.LocationId, 20)

	r1 := events.NewVirtualGroupCreated(id(), prod(), "Writer", []events.ProductLine{
		line(p1.ProductId, 2),
		line(p2.ProductId, 1),
	})

	return &module.UseCase{
		Name: "First test this is",
		Given: []shared.Event{
			p1, p2, l1, i1, i2, r1,
		},
		When: spec.Get("/reports/groups"),

		ThenResponse: spec.ReturnJSON(seq.Map{
			"items[0]": seq.Map{
				"groupId":  r1.GroupId,
				"quantity": 5,
				"name":     "Writer",
			},
		}),
		ThenEvents: spec.Events(),
	}
}

// syntax sugar
func id() shared.Id {
	return shared.NewId()
}
func loc() events.LocationId {
	return events.NewLocationId()
}

func prod() events.ProductId {
	return events.NewProductId()
}
func line(id events.ProductId, quantity int) events.ProductLine {
	return events.ProductLine{id, quantity}
}
