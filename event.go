package core

type Event interface {
	Meta() *Info
}

type Contract string

type Info struct {
	Contract string
	EventId  Id
}

func NewInfo(contract string, eventId Id) *Info {
	return &Info{contract, eventId}
}
