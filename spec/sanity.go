package spec

import (
	"fmt"
	"strconv"
)

func checkSanity(scenarios []*scenario) []string {
	out := []string{}

	bad := func(line string, args ...interface{}) {
		out = append(out, fmt.Sprintf(line, args...))
	}

	if len(scenarios) == 0 {
		bad("Nothing to verify. Did you add usecases?")
	}

	for i, s := range scenarios {
		uc := s.UseCase
		name := uc.Name
		if uc.Name == "" {
			name := strconv.Itoa(i)
			bad("Must have a name '%s'", name)
		}
		if len(uc.Given) == 0 && uc.When == nil {
			bad("Must have either given or when: '%s'", name)
		}
		if (uc.When == nil) != (uc.ThenResponse == nil) {
			bad("When and ThenResponse must be provided both: '%s'", name)
		}

	}
	return out
}
