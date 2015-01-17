package reports

import (
	"bitbucket.org/abdullin/proto/back/events"
	"bitbucket.org/abdullin/proto/back/seq"
	"bitbucket.org/abdullin/proto/back/shared"
	"bitbucket.org/abdullin/proto/back/spec"
)

func First() *spec.UseCase {
	p1 := events.NewProductCreated(id(), prod(), "pencil")
	p2 := events.NewProductCreated(id(), prod(), "notepad")

	l1 := events.NewLocationCreated(id(), loc(), "loc1")

	i1 := events.NewItemAdded(id(), p1.ProductId, l1.LocationId, 10)
	i2 := events.NewItemAdded(id(), p2.ProductId, l1.LocationId, 20)

	r1 := events.NewVirtualGroupCreated(id(), prod(), "Writer", events.ProductList{
		p1.ProductId: 2,
		p2.ProductId: 1,
	})

	return &spec.UseCase{
		Given: []shared.Event{
			p1, p2, l1, i1, i2, r1,
		},
		When: spec.Get("/reports/groups"),
		ThenResponse: seq.Map{
			"length": 1,
			"[0]": seq.Map{
				"groupId":  r1.GroupId,
				"quantity": 5,
				"name":     "Writer",
			},
		},
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
