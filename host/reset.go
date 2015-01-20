package main

import (
	"bitbucket.org/abdullin/proto/back/db"
	"bitbucket.org/abdullin/proto/back/host/setup"
	"github.com/op/go-logging"
)

func main() {
	var l = logging.MustGetLogger("reset")

	for _, s := range setup.Specs {
		if s.Schema == "" {
			continue
		}
		l.Info("Reset DB for %s", s.Name)

		db.CreateOrResetDB(s.Name, s.Name, s.Schema)

	}
}
