package shared

type Event interface {
	Meta() *Info
}

type Contract string

type Info struct {
	Contract Contract
	EventId  Id
}
