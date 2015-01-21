package spec

import (
	"github.com/abdullin/omni/core"
	"github.com/abdullin/omni/core/env"
)

type Values map[string]string

func GetJSON(url string, values Values) *env.Request {
	return &env.Request{"GET", url, nil, ""}
}

func GivenEvents(es ...core.Event) []core.Event {
	if len(es) == 0 {
		return []core.Event{}
	}
	return es
}
