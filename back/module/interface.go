package module

import "bitbucket.org/abdullin/proto/back/api"

type Registrar interface {
	HandleHttp(method string, path string, handler api.Handler)
}

type Module interface {
	Register(r Registrar)
}

type Factory func() Module

type Spec struct {
	Name    string
	Schema  string
	Factory Factory
}
