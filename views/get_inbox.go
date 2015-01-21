package views

import "github.com/abdullin/omni/core/api"

func (m *Module) getInbox(req *api.Request) api.Response {

	type task struct{}
	type response struct {
		Tasks []task `json:"tasks"`
	}

	return api.NewJSON(response{
		Tasks: []task{},
	})

}
