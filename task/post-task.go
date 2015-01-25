package task

import (
	"github.com/abdullin/omni/core/api"
	"github.com/abdullin/omni/lang"
)

type postTaskRequest struct {
	Name  string `json:"name"`
	Inbox bool   `json:"inbox"`
}
type postTaskResponse struct {
	TaskId lang.TaskId `json:"taskId"`
	Name   string      `json:"name"`
	Inbox  bool        `json:"inbox"`
}

func (m *Module) postTask(req *api.Request) api.Response {
	var request postTaskRequest

	if err := req.ParseBody(&request); err != nil {
		return api.BadRequestResponse(err.Error())
	}
	if request.Name == "" {
		return api.BadRequestResponse("Expected name")
	}

	// TODO - accept array
	taskId := lang.NewTaskId()
	m.pub.MustPublish(lang.NewTaskAdded(lang.NewEventId(), &taskId, request.Name))
	if request.Inbox {
		m.pub.MustPublish(lang.NewTaskMovedToInbox(lang.NewEventId(), taskId))
	}

	// track request ID for idempotency
	return api.NewJSON(&postTaskResponse{
		TaskId: taskId,
		Name:   request.Name,
		Inbox:  request.Inbox,
	})
}
