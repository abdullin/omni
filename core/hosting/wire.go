package hosting

import "github.com/abdullin/omni/core/env"

func New(pub env.Publisher, specs []*env.Spec) *Context {

	var c = &Context{}
	for _, spec := range specs {
		r := env.NewContainer()
		var m = spec.Factory(pub)
		m.Register(r)
		c.Items = append(c.Items, r)
	}
	return c

}
