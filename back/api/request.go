package api

import "net/http"

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
