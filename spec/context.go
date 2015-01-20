package spec

import "bitbucket.org/abdullin/proto/back/module"

func NewContext(spec *module.Spec) *Context {
	return &Context{
		pub:  newPublisher(),
		spec: spec,
	}
}

type Context struct {
	pub  *publisher
	spec *module.Spec
}

func (c *Context) Pub() module.Publisher {
	return c.pub
}

func (c *Context) Verify(m module.Module) *Report {
	return buildAndVerify(c.pub, c.spec, m)
}

func Get(url string) *module.Request {
	return &module.Request{"GET", url, nil, ""}
}
