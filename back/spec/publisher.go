package spec

import "bitbucket.org/abdullin/proto/back/shared"

func newPublisher() *publisher {
	return &publisher{}
}

type publisher struct {
	Events []shared.Event
}

func (p *publisher) MustPublish(e shared.Event) {
	if err := p.Publish(e); err != nil {
		panic(err)
	}
}
func (p *publisher) Publish(e shared.Event) error {
	p.Events = append(p.Events, e)
	return nil
}

func (p *publisher) Clear() {
	p.Events = nil
}
