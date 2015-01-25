package task

import (
	"github.com/abdullin/omni/core/env"
	"github.com/abdullin/omni/core/spec"
	"github.com/abdullin/omni/lang"
	"github.com/abdullin/seq"
)

var useCases = []env.UseCaseFactory{
	when_post_task_then_event_is_published,
}

func when_post_task_then_event_is_published() *env.UseCase {

	ignoreEventId := lang.NewEventId()
	newTaskId := lang.NewTaskId()

	return &env.UseCase{
		Name: "When POST /task to inbox, then 2 events are published",
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
			lang.NewTaskAdded(ignoreEventId, newTaskId, "NewTask"),
			lang.NewTaskMovedToInbox(ignoreEventId, newTaskId),
		),
		Where: spec.Where{
			newTaskId:     "newTaskId",
			ignoreEventId: "ignore",
		},
	}
}
