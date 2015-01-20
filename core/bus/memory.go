package bus

import (
	"errors"
	"sync"
	"time"

	"github.com/abdullin/omni/core"
	"github.com/abdullin/omni/core/env"
)

type mem struct {
	c        chan core.Event
	handlers map[string]env.EventHandler
	started  bool

	mu sync.Mutex
}

func (m *mem) AddEventHandler(name string, h env.EventHandler) {
	m.handlers[name] = h
}

func (m *mem) MustPublish(e core.Event) {
	m.c <- e
}
func (m *mem) Publish(e core.Event) error {
	m.c <- e
	return nil
}

func (m *mem) Start() {
	if m.started {
		panic("Bus can't be started twice")
	}

	m.started = true

	go func() {
		for {
			select {
			case message := <-m.c:
				m.dispatch(message)
			}
		}
	}()
}

func (m *mem) dispatch(e core.Event) {
	for name, h := range m.handlers {
		if err := handleWithTimeout(h, e); err != nil {
			var contract = e.Meta().Contract
			l.Panicf("%s @ %s: %s", name, contract, err)
		}
	}
}

func handleWithTimeout(h env.EventHandler, e core.Event) (err error) {
	c := make(chan error, 1)

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	go func() {
		c <- h.HandleEvent(e)
	}()

	select {
	case err = <-c:
		return

	case <-time.After(time.Second):
		err = errors.New("Timeout")
	}
	return
}

type Bus interface {
	AddEventHandler(name string, h env.EventHandler)
	env.Publisher
	Start()
}

func NewMem() Bus {
	return &mem{
		c:        make(chan core.Event, 10000),
		handlers: make(map[string]env.EventHandler),
	}
}
