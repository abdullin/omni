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

func i(contract string, eventId EventId) *Info {
	return NewInfo(contract, eventId.Id)
}

type TaskAdded struct {
	EventId EventId `json:"eventId"`
	TaskId  TaskId  `json:"taskId"`
	Name    string  `json:"name"`
	Inbox   bool    `json:"inbox"`
}

func NewTaskAdded(event EventId, task TaskId, name string, inbox bool) *TaskAdded {
	return &TaskAdded{event, task, name, inbox}
}

type TaskChecked struct {
	EventId EventId `json:"eventId"`
	TaskId  TaskId  `json:"taskId"`
}

func NewTaskChecked(event EventId, task TaskId) *TaskChecked {
	return &TaskChecked{event, task}
}

type TaskUnchecked struct {
	EventId EventId `json:"eventId"`
	TaskId  TaskId  `json:"taskId"`
}

func NewTaskUnhecked(event EventId, task TaskId) *TaskUnchecked {
	return &TaskUnchecked{event, task}
}

type TaskStarred struct {
	EventId EventId `json:"eventId"`
	TaskId  TaskId  `json:"taskId"`
}

func NewTaskStarred(event EventId, task TaskId) *TaskStarred {
	return &TaskStarred{event, task}
}

type TaskUnstarred struct {
	EventId EventId `json:"eventId"`
	TaskId  TaskId  `json:"taskId"`
}

func NewTaskUnstarred(event EventId, task TaskId) *TaskUnstarred {
	return &TaskUnstarred{event, task}
}

func (e *TaskAdded) Meta() *Info     { return i("TaskAdded", e.EventId) }
func (e *TaskChecked) Meta() *Info   { return i("TaskChecked", e.EventId) }
func (e *TaskStarred) Meta() *Info   { return i("TaskStarred", e.EventId) }
func (e *TaskUnstarred) Meta() *Info { return i("TaskUnstarred", e.EventId) }
