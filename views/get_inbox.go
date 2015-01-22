package views

import (
	"fmt"

	"github.com/abdullin/omni/core/api"
	"github.com/abdullin/omni/lang"
)

func (m *Module) getInbox(req *api.Request) api.Response {

	type task struct {
		TaskId lang.TaskId `json:"taskId"`
		Name   string      `json:"name"`
	}
	type response struct {
		Tasks []task `json:"tasks"`
	}

	var items = []task{}
	for k, t := range m.s.all {
		fmt.Println("Serve", k, t)
		items = append(items, task{t.TaskId, t.Name})
	}

	return api.NewJSON(response{
		Tasks: items,
	})
}
