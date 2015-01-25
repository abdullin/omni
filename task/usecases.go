package task

import (
	"github.com/abdullin/omni/core/env"
	"github.com/abdullin/omni/core/spec"
	"github.com/abdullin/omni/lang"
	"github.com/abdullin/seq"
)

var useCases = []env.UseCaseFactory{
	given_unchecked_task_when_check_then_event,
	when_post_inbox_task_then_event_is_published,
}

func newEventId() lang.EventId {
	return lang.NewEventId()
}
func newTaskId() lang.TaskId {
	return lang.NewTaskId()
}

var IgnoreEventId lang.EventId

func given_unchecked_task_when_check_then_event() *env.UseCase {

	taskId := lang.NewTaskId()

	return &env.UseCase{
		Name: "Given new task, when PUT /task with check, then event",
		Given: spec.Events(
			lang.NewTaskAdded(newEventId(), taskId, "ho-ho"),
		),
		When: spec.PutJSON("/task", seq.Map{
			"checked": true,
			"taskId":  taskId,
		}),
		ThenResponse: spec.ReturnJSON(seq.Map{
			"taskId":  taskId,
			"name":    "ho-ho",
			"checked": true,
			"starred": false,
		}),
		ThenEvents: spec.Events(
			lang.NewTaskChecked(IgnoreEventId, taskId),
		),
		Where: spec.Where{IgnoreEventId: "ignore"},
	}
}

func when_post_inbox_task_then_event_is_published() *env.UseCase {

	newTaskId := lang.NewTaskId()

	return &env.UseCase{
		Name: "When POST /task for inbox, then 2 events are published",
		When: spec.PostJSON("/task", seq.Map{
			"name":  "NewTask",
			"inbox": true,
		}),
		ThenResponse: spec.ReturnJSON(seq.Map{
			"name":   "NewTask",
			"inbox":  "true",
			"taskId": newTaskId,
		}),
		ThenEvents: spec.Events(
			lang.NewTaskAdded(IgnoreEventId, newTaskId, "NewTask"),
			lang.NewTaskMovedToInbox(IgnoreEventId, newTaskId),
		),
		Where: spec.Where{
			newTaskId:     "sameTaskId",
			IgnoreEventId: "ignore",
		},
	}
}
