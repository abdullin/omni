package lang

import . "github.com/abdullin/omni/core"

type TaskId struct{ Id }

func NewTaskId() TaskId {
	return TaskId{NewId()}
}

func i(contract string, eventId Id) *Info {
	return NewInfo(contract, eventId)
}

type TaskAdded struct {
	EventId Id     `json:"eventId"`
	TaskId  TaskId `json:"taskId"`
	Name    string `json:"name"`
	Inbox   bool   `json:"inbox"`
}

func NewTaskAdded(event Id, task TaskId, name string, inbox bool) *TaskAdded {
	return &TaskAdded{event, task, name, inbox}
}

type TaskChecked struct {
	EventId Id     `json:"eventId"`
	TaskId  TaskId `json:"taskId"`
}

func NewTaskChecked(event Id, task TaskId) *TaskChecked {
	return &TaskChecked{event, task}
}

type TaskUnchecked struct {
	EventId Id     `json:"eventId"`
	TaskId  TaskId `json:"taskId"`
}

func NewTaskUnhecked(event Id, task TaskId) *TaskUnchecked {
	return &TaskUnchecked{event, task}
}

func (e *TaskAdded) Meta() *Info   { return i("TaskAdded", e.EventId) }
func (e *TaskChecked) Meta() *Info { return i("TaskChecked", e.EventId) }

type TaskStarred struct {
	EventId Id     `json:"eventId"`
	TaskId  TaskId `json:"taskId"`
}

func NewTaskStarred(event Id, task TaskId) *TaskStarred {
	return &TaskStarred{event, task}
}

func (e *TaskStarred) Meta() *Info { return i("TaskStarred", e.EventId) }

type TaskUnstarred struct {
	EventId Id     `json:"eventId"`
	TaskId  TaskId `json:"taskId"`
}

func NewTaskUnstarred(event Id, task TaskId) *TaskUnstarred {
	return &TaskUnstarred{event, task}
}

func (e *TaskUnstarred) Meta() *Info { return i("TaskUnstarred", e.EventId) }
