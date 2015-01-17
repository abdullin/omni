package setup

func Modules() *Context {
	var c = &Context{}
	for _, spec := range Specs {
		r := newRegistrar()
		var m = spec.Factory()
		m.Register(r)
		c.Items = append(c.Items, r)
	}
	return c

}
