package spec

import "fmt"

func af(in []string, line string, args ...interface{}) []string {
	return append(in, fmt.Sprintf(line, args...))
}

func checkPreRunSanity(scenarios []*scenario) []string {
	out := []string{}

	if len(scenarios) == 0 {
		out = af(out, "Nothing to verify. Did you add usecases?")
	}
	return out
}
