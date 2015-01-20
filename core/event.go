package core

type Event interface {
	Meta() *Info
}

type Contract string

type Info struct {
	Contract string
	EventId  Id
}
