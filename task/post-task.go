package task

import (
	"github.com/abdullin/omni/core/api"
	"github.com/abdullin/omni/lang"
)

func (m *Module) postTask(req *api.Request) api.Response {

	m.pub.MustPublish(lang.NewTaskAdded(lang.NewEventId(), lang.NewTaskId(), "nil"))

	// track request ID for idempotency
	return api.NewError("Not implemented", 200)
}
