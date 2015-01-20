package spec

import (
	"net/http"

	"github.com/abdullin/omni/core/env"
)

// default errors are in JSON
func ReturnErrorJSON(status int) *module.Response {
	return &module.Response{
		Status:  status,
		Body:    nil,
		Headers: http.Header{},
	}
}

func ReturnJSON(response interface{}) *module.Response {
	return &module.Response{
		Status: http.StatusOK,
		Body:   response,

		Headers: http.Header{
			"Content-Type": []string{"application/json"},
		},
	}
}
