package spec

import (
	"net/http"

	"bitbucket.org/abdullin/proto/back/module"
)

func ReturnError(status int) *module.Response {
	return &module.Response{
		Status:  status,
		Body:    nil,
		Headers: http.Header{},
	}
}
