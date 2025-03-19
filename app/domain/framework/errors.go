package framework

import (
	"net/http"
)

var (
	ErrBadRequest          = NewControllerError(http.StatusBadRequest, "", nil)
	ErrNotFound            = NewControllerError(http.StatusNotFound, "", nil)
	ErrInternalServerError = NewControllerError(http.StatusInternalServerError, "", nil)
)

type ControllerError struct {
	Status  int
	Message string
	Err     error
}

func NewControllerError(status int, message string, err error) *ControllerError {
	return &ControllerError{
		Status:  status,
		Message: message,
		Err:     err,
	}
}

func (e *ControllerError) Error() string {
	v := e.Message
	if e.Err != nil {
		v += ": " + e.Err.Error()
	}
	return v
}

func (e *ControllerError) Unwrap() error { return e.Err }

func (e *ControllerError) With(msg string) *ControllerError {
	return &ControllerError{
		Status:  e.Status,
		Message: msg,
		Err:     e.Err,
	}
}

func (e *ControllerError) WithErr(err error) *ControllerError {
	return &ControllerError{
		Status:  e.Status,
		Message: e.Message,
		Err:     err,
	}
}
