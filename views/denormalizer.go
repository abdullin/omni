package views

import (
	"github.com/abdullin/omni/core"
	"github.com/abdullin/omni/lang"
)

type denormalizer struct {
	s *store
}

func newDenormalizer(s *store) *denormalizer {
	return &denormalizer{s}
}

func (d *denormalizer) HandleEvent(e core.Event) error {
	switch t := e.(type) {
	case *lang.TaskAdded:
		d.s.addTaskToInbox(t.TaskId, t.Name, t.Inbox)
	}
	return nil
}

func (d *denormalizer) Contracts() []string {
	return []string{
		"TaskAdded",
		"TaskCreated",
	}
}
