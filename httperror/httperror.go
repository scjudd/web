package httperror

import (
	"fmt"
	"net/http"
)

type Error interface {
	error
	Code() int
	HttpError(http.ResponseWriter)
}

func New(code int, err error) *httpError {
	return &httpError{err, code}
}

type httpError struct {
	error
	code int
}

func (err *httpError) Code() int {
	return err.code
}

func (err *httpError) HttpError(w http.ResponseWriter) {
	http.Error(w, fmt.Sprintf("%d %s", err.code, http.StatusText(err.code)), err.code)
}
