package bus

import (
	"fmt"

	"github.com/abdullin/omni"
	"github.com/abdullin/omni/env"
	"github.com/op/go-logging"
)

var (
	l = logging.MustGetLogger("bus")
)

type LoggingWrapper struct {
	inner env.Publisher
}

func WrapWithLogging(p env.Publisher) env.Publisher {
	return &LoggingWrapper{p}
}

func log(es ...core.Event) {
	for _, e := range es {
		switch e := e.(type) {
		case fmt.Stringer:
			l.Debug("%v", e.String())
		default:
			var contract = e.Meta().Contract
			l.Debug("%v", contract)
		}
	}
}

func (p *LoggingWrapper) Publish(e ...core.Event) error {
	log(e...)
	return p.inner.Publish(e...)
}
func (p *LoggingWrapper) MustPublish(e ...core.Event) {
	log(e...)
	p.inner.MustPublish(e...)
}
