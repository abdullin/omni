package task

import (
	"github.com/abdullin/omni/core/env"
	"github.com/abdullin/omni/core/spec"
	"github.com/abdullin/seq"
)

var useCases = []env.UseCaseFactory{
	when_post_task_then_event_is_published,
}

func when_post_task_then_event_is_published() *env.UseCase {
	return &env.UseCase{
		Name: "When POST /task, then event is published",
		When: spec.PostJSON("/task", seq.Map{
			"name":  "NewTask",
			"inbox": true,
		}),
		ThenResponse: spec.ReturnJSON(seq.Map{
			"name":  "NewTask",
			"inbox": "true",
		}),
		ThenEvents: spec.Events(
			seq.Map{
				"name": "NewTask",
			},
			seq.Map{},
		),
	}
}
