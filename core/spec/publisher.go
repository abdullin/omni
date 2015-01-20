package spec

import "github.com/abdullin/omni/core"

func newPublisher() *publisher {
	return &publisher{}
}

type publisher struct {
	Events []core.Event
}

func (p *publisher) MustPublish(e core.Event) {
	if err := p.Publish(e); err != nil {
		panic(err)
	}
}
func (p *publisher) Publish(e core.Event) error {
	p.Events = append(p.Events, e)
	return nil
}

func (p *publisher) Clear() {
	p.Events = nil
}
