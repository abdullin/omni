package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/abdullin/omni/core/bus"
	"github.com/abdullin/omni/core/env"
	"github.com/abdullin/omni/reports"

	"github.com/abdullin/omni/core/hosting"

	"github.com/op/go-logging"
)

var l = logging.MustGetLogger("main")

var specs = []*env.Spec{
	reports.Spec,
}

func main() {

	router := mux.NewRouter()
	memBus := bus.NewMem()

	var wrap = bus.WrapWithLogging(memBus)
	var host = hosting.New(wrap, specs)

	host.WireHttp(router)
	host.WireHandlers(memBus)

	memBus.Start()

	bind := ":8001"
	l.Info("Listening at %v", bind)
	l.Panic(http.ListenAndServe(bind, router))
}
