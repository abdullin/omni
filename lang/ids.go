package lang

import . "github.com/abdullin/omni/core"

type TaskId struct{ Id }
type EventId struct{ Id }

func NewTaskId() TaskId {
	return TaskId{NewId()}
}

func NewEventId() EventId {
	return EventId{NewId()}
}
