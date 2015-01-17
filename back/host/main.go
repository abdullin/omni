package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"bitbucket.org/abdullin/proto/back/host/setup"
	"github.com/op/go-logging"
)

var l = logging.MustGetLogger("main")

func main() {
	var ctx = setup.Modules()

	bind := ":8001"
	l.Info("Listening at %v", bind)

	router := mux.NewRouter()
	ctx.WireHttp(router)

	l.Panic(http.ListenAndServe(bind, router))
}
