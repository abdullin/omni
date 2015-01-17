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

	return &spec.UseCase{
		Given: []shared.Event{
			p1, p2, l1, i1, i2,
		},
		ThenEvents: seq.Map{
			"[0]": seq.Map{
				"contract": "ProductCreated",
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
