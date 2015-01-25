package lang

import . "github.com/abdullin/omni/core"

func i(contract string, eventId EventId) *Info {
	return NewInfo(contract, eventId.Id)
}

type TaskAdded struct {
	EventId EventId `json:"eventId,omitempty"`
	TaskId  *TaskId `json:"taskId,omitempty"`
	Name    string  `json:"name"`
}

func NewTaskAdded(event EventId, task *TaskId, name string) *TaskAdded {
	return &TaskAdded{event, task, name}
}

type TaskMovedToInbox struct {
	EventId EventId `json:"eventId"`
	TaskId  TaskId  `json:"taskId"`
}

func NewTaskMovedToInbox(event EventId, task TaskId) *TaskMovedToInbox {
	return &TaskMovedToInbox{event, task}
}

func (e *TaskMovedToInbox) Meta() *Info { return i("TaskMovedToInbox", e.EventId) }

type TaskRemoved struct {
	EventId EventId `json:"eventId"`
	TaskId  TaskId  `json:"taskId"`
}

func NewTaskRemoved(event EventId, task TaskId) *TaskRemoved {
	return &TaskRemoved{event, task}
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

type ContextCreated struct {
	EventId   EventId   `json:"eventId"`
	ContextId ContextId `json:"contextId"`
	Name      string    `json:"name"`
}

func NewContextCreated(event EventId, context ContextId, name string) *ContextCreated {
	return &ContextCreated{event, context, name}
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

func (e *TaskRemoved) Meta() *Info { return i("TaskRemoved", e.EventId) }

func (e *ContextCreated) Meta() *Info { return i("ContextCreated", e.EventId) }
