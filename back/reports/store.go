package reports

import "bitbucket.org/abdullin/proto/back/shared"

type store struct {
	Events []shared.Event
}

func (s *store) AddEvent(e shared.Event) {
	s.Events = append(s.Events, e)
}
