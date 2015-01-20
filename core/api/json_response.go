package api

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	subj interface{}
	Code int
}

func NewJSON(subj interface{}) Response {
	return &jsonResponse{subj, 200}
}

func (e *jsonResponse) Write(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(e.subj)
	guard("marhsal", err)
	w.WriteHeader(e.Code)
	w.Write(b)
}
