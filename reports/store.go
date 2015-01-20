package reports

import "github.com/abdullin/omni/shared"

type store struct {
	Events []shared.Event
}

func (s *store) AddEvent(e shared.Event) {
	s.Events = append(s.Events, e)
}
