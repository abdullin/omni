package reports

import "github.com/abdullin/omni/core"

type store struct {
	Events []shared.Event
}

func (s *store) AddEvent(e shared.Event) {
	s.Events = append(s.Events, e)
}
