package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Request struct {
	Raw *http.Request
}

func NewRequest(inner *http.Request) *Request {
	return &Request{inner}
}

func (r *Request) String(param string) string {
	if v := r.Raw.FormValue(param); v == "" {
		panic(NewErrArgumentNil(param))
	} else {
		return v
	}
}

func (r *Request) ParseBody(subj interface{}) error {
	content := r.Raw.Header.Get("Content-Type")
	switch content {
	case "application/json":
		decoder := json.NewDecoder(r.Raw.Body)
		return decoder.Decode(subj)

	}
	return errors.New("Unexpected content type " + content)

}
