package spec

import (
	"fmt"
	"testing"
)

type Results struct {
	Insanity []string
}

func NewResults() *Results {
	return &Results{}
}

func (r *Results) failSanity(s string, args ...interface{}) {
	r.Insanity = append(r.Insanity, fmt.Sprintf(s, args...))
}

func (r *Results) Report(t *testing.T) {

	var fail = len(r.Insanity)

	if fail > 0 {
		t.Fail()
	}

}
