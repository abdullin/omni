package api

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func handleError(w http.ResponseWriter, err error) {
	switch err := err.(type) {
	case *ErrArgument:
		NewError(err.Error(), http.StatusBadRequest).Write(w)
	default:
		NewError(err.Error(), http.StatusInternalServerError).Write(w)
		l.Error("ERR : %v.\n%s", err.Error(), debug.Stack())
	}
}

func NewError(error string, code int) Response {
	return &errorResponse{code, error}
}

func BadRequestResponse(err string) Response {
	return &errorResponse{
		Code:  http.StatusBadRequest,
		Error: err,
	}
}

func NotImplementedResponse() Response {
	return &errorResponse{
		Code:  http.StatusInternalServerError,
		Error: "Not implemented",
	}
}

type ErrArgument struct {
	Argument string
	Problem  string
}

func (e *ErrArgument) Error() string {
	return fmt.Sprintf("Argument '%v': %v", e.Argument, e.Problem)
}

func NewErrArgumentNil(param string) *ErrArgument {
	return &ErrArgument{param, "can't be empty"}
}

func NewErrArgument(param, problem string) *ErrArgument {
	return &ErrArgument{param, problem}
}
