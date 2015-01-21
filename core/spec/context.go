package spec

import "github.com/abdullin/omni/core/env"

func NewContext(spec *env.Spec) *Context {
	return &Context{
		pub:  newPublisher(),
		spec: spec,
	}
}

type Context struct {
	pub  *publisher
	spec *env.Spec
}

func (c *Context) Pub() env.Publisher {
	return c.pub
}

func (c *Context) Verify(m env.Module) *Report {
	return buildAndVerify(c.pub, c.spec, m)
}
