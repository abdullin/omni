package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/op/go-logging"
)

var (
	l = logging.MustGetLogger("api")
)

func guard(name string, err error) {
	if err != nil {
		l.Panicf("%s: %s", name, err)
	}
}

func Handle(router *mux.Router, route *Route) *mux.Route {
	var handler = func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					handleError(w, err)
					return
				}
				panic(r)
			}
		}()

		// TODO: setup request
		var request = NewRequest(req)
		var response = route.Handler(request)
		response.Write(w)
	}

	return router.Methods(route.Method).Subrouter().HandleFunc(route.Path, handler)
}
