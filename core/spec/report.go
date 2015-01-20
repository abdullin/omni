package spec

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abdullin/omni/core/env"
	"github.com/abdullin/omni/core"
)

// Scenario result
type Result struct {
	UseCase  *module.UseCase
	Events   []shared.Event
	Response *httptest.ResponseRecorder

	EventsDiffs   []string
	ResponseDiffs []string
}

func (r *Result) Ok() bool {
	return len(r.ResponseDiffs)+len(r.EventsDiffs) == 0
}

type Report struct {
	Insanity []string
	Resuls   []*Result
}

func NewReport() *Report {
	return &Report{}
}

func (r *Report) failSanity(s string, args ...interface{}) {
	r.Insanity = append(r.Insanity, fmt.Sprintf(s, args...))
}

func prettyPrintEvent(e shared.Event) string {
	if e == nil {
		return fmt.Sprintf("%T<nil>", e)
	}
	if s, ok := e.(fmt.Stringer); ok {
		return fmt.Sprintf("%s", s.String())
	} else {

		return fmt.Sprintf("%s: %s", e.Meta().Contract, string(marshal(e)))
	}
}

func (r *Report) ToTesting(t *testing.T) {

	for _, x := range r.Insanity {
		t.Fail()
		fmt.Println("☹", x)
	}

	for _, r := range r.Resuls {
		if !r.Ok() {
			t.Fail()
			fmt.Println("X", r.UseCase.Name)

			if len(r.UseCase.Given) > 0 {
				fmt.Println("GIVEN")
				for i, e := range r.UseCase.Given {
					fmt.Println("  ", i+1, prettyPrintEvent(e))

				}
			}

			if r.UseCase.When != nil {
				when := r.UseCase.When
				uri, err := url.Parse(when.Path)
				guard("url.Parse", err)
				fmt.Println("WHEN", when.Method, uri.Path)
				query := uri.Query()
				if len(query) > 0 {
					fmt.Println("  with")

					for k, _ := range query {
						fmt.Println("  %s = '%s'", k, query.Get(k))
					}

				}
			}

			if resp := r.UseCase.ThenResponse; resp != nil {
				fmt.Println("EXPECT HTTP", resp.Status)
				if resp.Body != nil {
					fmt.Println(string(marshalIndent(resp.Body)))
				}
			}

			if es := r.UseCase.ThenEvents; es != nil {
				fmt.Println("EXPECT", len(es), "event(s)")
			}

			if len(r.ResponseDiffs) > 0 {

				fmt.Println("Response issues:")
				for _, x := range r.ResponseDiffs {
					fmt.Println("  " + x)
				}
			}

			if len(r.EventsDiffs) > 0 {
				fmt.Println("Event issues:")
				for _, x := range r.EventsDiffs {
					fmt.Println("  " + x)
				}
			}

		} else {
			fmt.Println("V", r.UseCase.Name)
		}
	}

}