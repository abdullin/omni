package spec

import (
	"fmt"
	"net/http/httptest"

	"bitbucket.org/abdullin/proto/back/api"
	"bitbucket.org/abdullin/proto/back/module"
	"github.com/gorilla/mux"
)

type scenario struct {
	UseCase *module.UseCase
}

func makeScenarios(ucs []module.UseCaseFactory) []*scenario {

	var out = []*scenario{}

	for _, f := range ucs {
		uc := f()
		sc := &scenario{uc}
		out = append(out, sc)

	}
	return out

}

func buildAndVerify(pub *publisher, spec *module.Spec, mod module.Module) *Results {

	var results = NewResults()

	results.failSanity("Spec tester not implemented")

	// TODO: sanity checks
	container := module.NewContainer()
	mod.Register(container)

	// wire routes

	scenarios := makeScenarios(spec.UseCases)
	router := mux.NewRouter()

	for _, route := range container.Routes {
		api.Handle(router, route)
	}

	for _, s := range scenarios {
		fmt.Println("Printing", s.UseCase.Name)
		// TODO: move to a separate method

		server := httptest.NewServer(router)

		for _, e := range s.UseCase.Given {
			contract := e.Meta().Contract
			for _, h := range container.Handlers {
				if needs(h, contract) {
					h.HandleEvent(e)
				}
			}
		}

		defer server.Close()

	}

	return results
}

func needs(eh module.EventHandler, contract string) bool {
	for _, c := range eh.Contracts() {
		if c == contract {
			return true
		}
	}
	return false
}
