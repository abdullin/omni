package spec

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"bytes"

	"github.com/abdullin/omni/core"
	"github.com/abdullin/omni/core/api"
	"github.com/abdullin/omni/core/env"
	"github.com/abdullin/seq"
	"github.com/gorilla/mux"
)

type scenario struct {
	UseCase *env.UseCase
}

func makeScenarios(ucs []env.UseCaseFactory) []*scenario {

	var out = []*scenario{}

	for _, f := range ucs {
		uc := f()
		sc := &scenario{uc}
		out = append(out, sc)

	}
	return out

}

func buildAndVerify(pub *publisher, spec *env.Spec, mod env.Module) *Report {

	var report = NewReport()

	// TODO: sanity checks
	container := env.NewContainer()
	mod.Register(container)
	scenarios := makeScenarios(spec.UseCases)

	insanities := checkSanity(scenarios)
	if len(insanities) > 0 {
		// fail fast
		report.addInsanities(insanities)
		return report
	}

	router := mux.NewRouter()

	// wire routes
	for _, route := range container.Routes {
		api.Handle(router, route)
	}

	for _, s := range scenarios {

		// reset data
		pub.Clear()
		for _, r := range container.DataReset {
			r()
		}

		result := &Result{
			UseCase: s.UseCase,
		}

		dispatchEvents(s.UseCase.Given, container.Handlers)

		if s.UseCase.When != nil {
			response := performRequest(s.UseCase.When, router)
			decodedResponse := decodeResponse(response)
			responseResult := verifyResponse(s.UseCase.ThenResponse, decodedResponse)

			result.ResponseDiffs = responseResult
			result.ResponseRaw = response
			result.Response = decodedResponse
		}
		{
			events := pub.Events
			eventsResult := verifyEvents(s.UseCase.ThenEvents, events)

			result.Events = events
			result.EventsDiffs = eventsResult
		}

		report.Resuls = append(report.Resuls, result)
	}

	return report
}

func guard(name string, err error) {
	if err != nil {
		panic(fmt.Errorf("%s: %s", name, err))
	}
}
func dispatchEvents(given []core.Event, handlers env.EventHandlerMap) {

	for _, e := range given {
		contract := e.Meta().Contract
		for _, h := range handlers {
			for _, c := range h.Contracts() {
				if c == contract {
					h.HandleEvent(e)
				}
			}
		}
	}
}

func performRequest(when *env.Request, router http.Handler) *httptest.ResponseRecorder {

	server := httptest.NewServer(router)
	defer server.Close()

	root := server.URL

	var buffer *bytes.Buffer

	if when.Body != nil {
		buffer = bytes.NewBuffer(marshal(when.Body))
	}

	request, err1 := http.NewRequest(when.Method, root+when.Path, buffer)
	guard("new request", err1)

	request.Header = when.Headers
	response, err2 := http.DefaultClient.Do(request)
	guard("response", err2)

	recorder := httptest.NewRecorder()

	if _, err := io.Copy(recorder, response.Body); err != nil {
		panic(err)
	}
	recorder.Code = response.StatusCode
	recorder.HeaderMap = response.Header

	return recorder
}

func decodeBody(response *httptest.ResponseRecorder) interface{} {
	if response.Body == nil {
		return nil
	}
	contentType := response.HeaderMap.Get("Content-Type")

	switch contentType {
	case "application/json":
		var body map[string]interface{}
		unmarshal(response.Body.Bytes(), &body)
		return body
	default:
		return response.Body.String()

	}
}

func decodeResponse(actual *httptest.ResponseRecorder) *env.Response {
	return &env.Response{
		Status:  actual.Code,
		Headers: actual.HeaderMap,
		Body:    decodeBody(actual),
	}
}

func verifyEvents(then []interface{}, actual []core.Event) []string {
	result := seq.Test(then, actual)
	return result.Diffs
}

func verifyResponse(then *env.Response, decoded *env.Response) []string {
	expected := seq.Map{
		"Status":  then.Status,
		"Headers": then.Headers,
		"Body":    then.Body,
	}

	//if then.Body != nil {
	//	expected["Body"] = then.Body
	//}

	actual := seq.Map{
		"Status":  decoded.Status,
		"Headers": decoded.Headers,
		"Body":    decoded.Body,
	}

	result := seq.Test(expected, actual)
	return result.Diffs

}
