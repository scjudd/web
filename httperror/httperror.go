package httperror

import (
	"fmt"
	"net/http"
)

type Error interface {
	error
	HttpError(http.ResponseWriter)
}

type Code interface {
	Code() int
}

func New(code int, err error) *httpError {
	return &httpError{code, err}
}

type httpError struct {
	code int
	err  error
}

func (err *httpError) Error() string {
	return err.err.Error()
}

func (err *httpError) HttpError(w http.ResponseWriter) {
	http.Error(w, fmt.Sprintf("%d %s", err.code, http.StatusText(err.code)), err.code)
}

func (err *httpError) Code() int {
	return err.code
}
