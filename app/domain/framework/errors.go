package framework

import (
	"fmt"
)

const sep = ": "

// BadRequestError is an error that http handler returns
// bad-request response to client on receive.
type BadRequestError struct {
	msg string
	err error
}

// NewBadRequestError returns new BadRequestError
func NewBadRequestError(msg string, err error) error {
	return &BadRequestError{msg: msg, err: err}
}
func (e *BadRequestError) Error() string {
	if e.err != nil {
		return e.msg + sep + e.err.Error()
	}
	return e.msg + sep
}
func (e *BadRequestError) Unwrap() error { return e.err }

// NotFoundError is an error that http handler returns
// not-found response to client on receive.
type NotFoundError struct {
	msg string
	err error
}

// NewNotFoundError returns new NotFoundError
func NewNotFoundError(msg string, err error) error {
	return &NotFoundError{msg: msg, err: err}
}
func (e *NotFoundError) Error() string {
	if e.err != nil {
		return e.msg + sep + e.err.Error()
	}
	return e.msg + sep
}
func (e *NotFoundError) Unwrap() error { return e.err }

// InternalServerError is an error that http handler returns
// internal-server-error response to client on receive.
type InternalServerError struct {
	msg string
	err error
}

// NewInternalServerError returns new InternalServerError
func NewInternalServerError(msg string, err error) error {
	return &InternalServerError{msg: msg, err: err}
}
func (e *InternalServerError) Error() string {
	if e.err != nil {
		return e.msg + sep + e.err.Error()
	}
	return e.msg + sep
}
func (e *InternalServerError) Unwrap() error { return e.err }

// // MustHandleError is an error that should be converted to other error type
// // before reach to http error handler.
// type MustHandleError struct {
// 	msg string
// 	err error
// }
//
// // NewMustHandleError returns new MustHandleError
// func NewMustHandleError(msg string, err error) error {
// 	return &MustHandleError{msg: msg, err: err}
// }
// func (e *MustHandleError) Error() string {
// 	if e.err != nil {
// 		return e.msg + sep + e.err.Error()
// 	}
// 	return e.msg + sep
// }
// func (e *MustHandleError) Unwrap() error { return e.err }

// WrapError wraps and returns error
func WrapError(msg string, err error) error {
	return fmt.Errorf(msg+sep+"%w", err)
}
