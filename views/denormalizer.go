package views

import "github.com/abdullin/omni/core"

type denormalizer struct {
	s *store
}

func newDenormalizer(s *store) *denormalizer {
	return &denormalizer{s}
}

func (d *denormalizer) HandleEvent(e core.Event) error {
	return nil
}

func (d *denormalizer) Contracts() []string {
	return []string{
		"TaskAdded",
		"TaskCreated",
	}
}
