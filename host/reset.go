package main

import (
	"github.com/abdullin/omni/db"
	"github.com/abdullin/omni/host/setup"
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
