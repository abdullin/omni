package hosting

import "github.com/abdullin/omni/env"

func New(modules []env.Module) *Context {
	context := &Context{}
	for _, mod := range modules {
		container := env.NewContainer()
		mod.Register(container)
		context.Items = append(context.Items, container)
	}
	return context

}
