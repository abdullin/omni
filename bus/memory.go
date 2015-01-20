package bus

import (
	"errors"
	"sync"
	"time"

	"github.com/abdullin/omni/module"
	"github.com/abdullin/omni/shared"
)

type mem struct {
	c        chan shared.Event
	handlers map[string]module.EventHandler
	started  bool

	mu sync.Mutex
}

func (m *mem) AddEventHandler(name string, h module.EventHandler) {
	m.handlers[name] = h
}

func (m *mem) MustPublish(e shared.Event) {
	m.c <- e
}
func (m *mem) Publish(e shared.Event) error {
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

func (m *mem) dispatch(e shared.Event) {
	for name, h := range m.handlers {
		if err := handleWithTimeout(h, e); err != nil {
			var contract = e.Meta().Contract
			l.Panicf("%s @ %s: %s", name, contract, err)
		}
	}
}

func handleWithTimeout(h module.EventHandler, e shared.Event) (err error) {
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
	AddEventHandler(name string, h module.EventHandler)
	module.Publisher
	Start()
}

func NewMem() Bus {
	return &mem{
		c:        make(chan shared.Event, 10000),
		handlers: make(map[string]module.EventHandler),
	}
}
