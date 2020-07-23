package httperror

import (
	"fmt"
	"net/http"
)

// Error implements error interface
type Error struct {
	message string
	code    int
	err     error
}

// UnauthorizedErr represents a 401
func UnauthorizedErr() Error {
	return Error{message: "unauthorized", code: http.StatusUnauthorized}
}

// NotFoundErr represents a 404
func NotFoundErr() Error {
	return Error{message: "not found", code: http.StatusNotFound}
}

// UnprocessableErr represents a 422
func UnprocessableErr() Error {
	return Error{message: "unable to process error", code: http.StatusUnprocessableEntity}
}

// Error returns the error message
func (s Error) Error() string {
	return fmt.Sprintf("%s : %s", s.message, s.err)
}

// StatusCode returns the http status code
func (s Error) StatusCode() int {
	return s.code
}

// SetMessage updates the message and returns a copy
func (s Error) SetMessage(msg string) Error {
	s.message = msg

	return s
}

// Wrap is useful for wrapping another error
func (s Error) Wrap(err error) Error {
	s.err = err

	return s
}

// Unwrap returns the wrapped error
func (s Error) Unwrap() error {
	return s.err
}

// New builds a httpError
func New(msg string, code int, err error) Error {
	return Error{msg, code, err}
}
