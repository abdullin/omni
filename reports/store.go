package reports

import "github.com/abdullin/omni/core"

type store struct {
	Events []core.Event
}

func (s *store) AddEvent(e core.Event) {
	s.Events = append(s.Events, e)
}
