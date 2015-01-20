package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/abdullin/omni/bus"
	"github.com/abdullin/omni/host/setup"
	"github.com/op/go-logging"
)

var l = logging.MustGetLogger("main")

func main() {

	router := mux.NewRouter()
	memBus := bus.NewMem()

	var wrap = bus.WrapWithLogging(memBus)

	var ctx = setup.Modules(wrap)

	bind := ":8001"
	l.Info("Listening at %v", bind)

	ctx.WireHttp(router)
	ctx.WireHandlers(memBus)

	memBus.Start()

	l.Panic(http.ListenAndServe(bind, router))
}
