package spec

import (
	"testing"

	"bitbucket.org/abdullin/proto/back/module"
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

func NewContext(spec *module.Spec) *Context {
	return &Context{}
}

type Context struct {
	Pub module.Publisher
}

func (c *Context) Verify(m module.Module) *Results {
	return nil
}

type Results struct {
}

func (v *Results) Report(t *testing.T) {}

type Request struct {
	Url    string
	Method string
}

func Get(url string) *Request {
	return &Request{url, "GET"}
}
