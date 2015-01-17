package reports

import (
	"bitbucket.org/abdullin/proto/back/events"
	"bitbucket.org/abdullin/proto/back/shared"
)

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
