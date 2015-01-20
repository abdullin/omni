package setup

import "bitbucket.org/abdullin/proto/back/module"

func Modules(pub module.Publisher) *Context {

	var c = &Context{}
	for _, spec := range Specs {
		r := newRegistrar()
		var m = spec.Factory(pub)
		m.Register(r)
		c.Items = append(c.Items, r)
	}
	return c

}
