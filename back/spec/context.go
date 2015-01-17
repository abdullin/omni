package spec

import (
	"bitbucket.org/abdullin/proto/back/seq"
	"bitbucket.org/abdullin/proto/back/shared"
)

type UseCase struct {
	Name string

	Given []shared.Event

	ThenEvents   seq.Map
	ThenResponse seq.Map
}

func NewContext() {}
