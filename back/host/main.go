package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"bitbucket.org/abdullin/proto/back/bus"
	"bitbucket.org/abdullin/proto/back/host/setup"
	"github.com/op/go-logging"
)

var l = logging.MustGetLogger("main")

func main() {

	router := mux.NewRouter()
	bus := bus.NewMem()

	var ctx = setup.Modules(bus)

	bind := ":8001"
	l.Info("Listening at %v", bind)

	ctx.WireHttp(router)
	ctx.WireHandlers(bus)

	bus.Start()

	l.Panic(http.ListenAndServe(bind, router))
}
