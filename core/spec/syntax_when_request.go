package spec

import (
	"net/http"

	"github.com/abdullin/omni/core"
	"github.com/abdullin/omni/core/env"
)

type Values map[string]string

func GetJSON(url string, values Values) *env.Request {
	return &env.Request{"GET", url, nil, ""}
}
func PostJSON(url string, subj interface{}) *env.Request {
	return &env.Request{
		Method: "POST",
		Path:   url,
		Headers: http.Header{
			"Content-Type": []string{"application/json"},
		},
		// headers
		Body: subj,
	}
}

func GivenEvents(es ...core.Event) []core.Event {
	if len(es) == 0 {
		return []core.Event{}
	}
	return es
}
