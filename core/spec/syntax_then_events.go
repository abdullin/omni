package spec

import "github.com/abdullin/omni/core"

func Events(events ...core.Event) []core.Event {
	if len(events) == 0 {
		return []core.Event{}
	}
	return events

}
