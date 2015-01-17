package spec

import (
	"bitbucket.org/abdullin/proto/back/seq"
	"bitbucket.org/abdullin/proto/back/shared"
)

type UseCase struct {
	Name string

	Given []shared.Event
	When  *Request

	ThenEvents   seq.Map
	ThenResponse seq.Map
}

func NewContext() {}

type Request struct {
	Url    string
	Method string
}

func Get(url string) *Request {
	return &Request{url, "GET"}
}
