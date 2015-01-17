package log

import (
	"os"

	"github.com/op/go-logging"
)

var l = logging.MustGetLogger("setup")

func Init(prefix string) {

	format := "%{time:15:04:05} %{level:.1s} (%{module}): %{message}"
	fmt := logging.MustStringFormatter(format)
	logging.SetFormatter(fmt)

	var backends []logging.Backend

	// Setup one stdout and one syslog backend.
	logBackend := logging.NewLogBackend(os.Stdout, "", 0)
	logBackend.Color = true
	backends = append(backends, logBackend)

	// also write to session
	if sys, err := logging.NewSyslogBackend(prefix); err != nil {
		l.Warning("syslog unavalable: %s", err)
	} else {
		backends = append(backends, sys)
	}

	logging.SetBackend(backends...)
}
