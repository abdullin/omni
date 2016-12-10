package spec

import "github.com/abdullin/omni/core"

func newPublisher() *publisher {
	return &publisher{}
}

type publisher struct {
	Events []core.Event
}

func (p *publisher) MustPublish(es ...core.Event) {
	if err := p.Publish(es...); err != nil {
		panic(err)
	}
}
func (p *publisher) Publish(es ...core.Event) error {
	p.Events = append(p.Events, es...)
	return nil
}

func (p *publisher) Clear() {
	p.Events = nil
}
