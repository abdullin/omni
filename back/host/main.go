package main

import (
	"net/http"

	"bitbucket.org/abdullin/proto/back/log"

	"github.com/gorilla/mux"

	"bitbucket.org/abdullin/proto/back/host/setup"
	"github.com/op/go-logging"
)

var l = logging.MustGetLogger("main")

func main() {
	log.Init("gtd")

	var ctx = setup.Modules()

	bind := ":8001"
	l.Info("Listening at %v", bind)

	router := mux.NewRouter()
	ctx.Wire(router)

	l.Panic(http.ListenAndServe(bind, router))

	l.Info("Running")
}
