package spec

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abdullin/omni/core"
	"github.com/abdullin/omni/core/env"
)

// Scenario result
type Result struct {
	UseCase     *env.UseCase
	Events      []core.Event
	ResponseRaw *httptest.ResponseRecorder
	Response    *env.Response

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

func (r *Report) addInsanities(insanities []string) {
	r.Insanity = append(r.Insanity, insanities...)
}

func prettyPrintEvent(e core.Event) string {
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

	fmt.Println("================= SUMMARY =================")
	for _, x := range r.Insanity {
		t.Fail()
		fmt.Println("☹", x)
	}

	specFailed := false
	for _, r := range r.Resuls {
		if r.Ok() {
			fmt.Println("♥", r.UseCase.Name)
		} else {
			specFailed = true

			t.Fail()
			fmt.Println("✗", r.UseCase.Name)
		}
	}

	if specFailed {
		fmt.Println("\n================= DETAILS =================")

		for _, r := range r.Resuls {
			if !r.Ok() {

				fmt.Println("✗", r.UseCase.Name)
				printDetail(r)
			}

		}
	}

}

func printDetail(r *Result) {

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
				fmt.Printf("  %s = '%s'\n", k, query.Get(k))
			}

		}
	}

	if resp := r.UseCase.ThenResponse; resp != nil {
		fmt.Printf("EXPECT ")
		printResponse(resp)
	}
	if resp := r.Response; resp != nil {
		fmt.Printf("ACTUAL ")
		printResponse(resp)
	}

	if es := r.UseCase.ThenEvents; es != nil {
		fmt.Println("EXPECT", len(es), "event(s)")

		for i, e := range es {
			fmt.Printf("%v. %s\n", i, string(marshal(e)))
		}
	}
	if es := r.Events; es != nil {
		fmt.Println("ACTUAL", len(es), "event(s)")
		for i, e := range es {
			fmt.Printf("%v. %s\n", i, string(marshal(e)))
		}
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
}

func printResponse(resp *env.Response) {

	body := ""

	if resp.Body != nil {
		body = (string(marshalIndent(resp.Body)))
	}
	fmt.Println("HTTP", resp.Status, body)
}
