package spec

import "bitbucket.org/abdullin/proto/back/shared"

func newPublisher() *publisher {
	return &publisher{}
}

type publisher struct {
	Events []shared.Event
}

func (this *publisher) MustPublish(e shared.Event) {
	if err := this.Publish(e); err != nil {
		panic(err)
	}
}
func (this *publisher) Publish(e shared.Event) error {
	this.Events = append(this.Events, e)
	return nil
}
