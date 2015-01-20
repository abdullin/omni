package bus

import (
	"fmt"

	"github.com/abdullin/omni/core/env"
	"github.com/abdullin/omni/core"
	"github.com/op/go-logging"
)

var (
	l = logging.MustGetLogger("bus")
)

type LoggingWrapper struct {
	inner module.Publisher
}

func WrapWithLogging(p module.Publisher) module.Publisher {
	return &LoggingWrapper{p}
}

func log(e shared.Event) {
	switch e := e.(type) {
	case fmt.Stringer:
		l.Debug("%v", e.String())
	default:
		var contract = e.Meta().Contract
		l.Debug("%v", contract)
	}
}

func (p *LoggingWrapper) Publish(e shared.Event) error {
	log(e)
	return p.inner.Publish(e)
}
func (p *LoggingWrapper) MustPublish(e shared.Event) {
	log(e)
	p.inner.MustPublish(e)
}
