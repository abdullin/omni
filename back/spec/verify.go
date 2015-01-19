package spec

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"bytes"

	"bitbucket.org/abdullin/proto/back/api"
	"bitbucket.org/abdullin/proto/back/module"
	"bitbucket.org/abdullin/proto/back/shared"
	"github.com/abdullin/seq"
	"github.com/gorilla/mux"
)

type scenario struct {
	UseCase *module.UseCase
}

func makeScenarios(ucs []module.UseCaseFactory) []*scenario {

	var out = []*scenario{}

	for _, f := range ucs {
		uc := f()
		sc := &scenario{uc}
		out = append(out, sc)

	}
	return out

}

func buildAndVerify(pub *publisher, spec *module.Spec, mod module.Module) *Report {

	var report = NewReport()

	// TODO: sanity checks
	container := module.NewContainer()
	mod.Register(container)

	// wire routes

	scenarios := makeScenarios(spec.UseCases)
	router := mux.NewRouter()

	for _, route := range container.Routes {
		api.Handle(router, route)
	}

	for _, s := range scenarios {

		// reset data
		pub.Clear()

		dispatchEvents(s.UseCase.Given, container.Handlers)

		response := performRequest(s.UseCase.When, router)
		events := pub.Events

		//fmt.Println(response)

		eventsResult := seq.Test(s.UseCase.ThenEvents, events)

		responseResult := verifyResponse(s.UseCase.ThenResponse, response)

		//fmt.Println(responseResult.Diffs)
		result := &Result{
			UseCase:       s.UseCase,
			Response:      response,
			Events:        events,
			EventsDiffs:   eventsResult.Diffs,
			ResponseDiffs: responseResult,
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
func dispatchEvents(given []shared.Event, handlers module.EventHandlerMap) {

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

func performRequest(when *module.Request, router http.Handler) *httptest.ResponseRecorder {
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
		var body interface{}
		unmarshal(response.Body.Bytes(), body)
		return body
	default:
		return response.Body.String()

	}

}

func verifyResponse(then *module.Response, recorded *httptest.ResponseRecorder) []string {
	expected := seq.Map{
		"Status":  then.Status,
		"Headers": then.Headers,
	}

	if then.Body != nil {
		expected["Body"] = then.Body
	}

	actual := seq.Map{
		"Status":  recorded.Code,
		"Headers": recorded.HeaderMap,
		"Body":    decodeBody(recorded),
	}

	result := seq.Test(expected, actual)
	return result.Diffs

}
