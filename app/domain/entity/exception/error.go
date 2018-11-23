// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exception

import (
	"fmt"
)

type Error struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s %s", e.Code, e.Msg)
}

func (e *Error) WithMsg(m string) *Error {
	return &Error{e.Code, m}
}

var (
	// http
	NotFound            = &Error{Code: "NotFound"}
	MethodNotAllowed    = &Error{Code: "MethodNotAllowed"}
	InternalServerError = &Error{Code: "InternalServerError"}
	FailedToReadBody    = &Error{Code: "FailedToReadBody"}
	FailedToParseJson   = &Error{Code: "FailedToParseJson"}

	// engine
	ConnectionTimeout      = &Error{Code: "ConnectionTimeout"}
	FailedToConnect        = &Error{Code: "FailedToConnect"}
	FailedToClose          = &Error{Code: "FailedToClose"}
	EngineIsNotRunning     = &Error{Code: "EngineIsNotRunning"}
	EngineIsAlreadyRunning = &Error{Code: "EngineIsAlreadyRunning"}
	FailedToExecCommand    = &Error{Code: "FailedToExecCommand"}
	FailedToStart          = &Error{Code: "FailedToStart"}
	FailedToUpdateOption   = &Error{Code: "FailedToUpdateOption"}
	InvalidOptionSyntax    = &Error{Code: "InvalidOptionSyntax"}
	InvalidOptionParameter = &Error{Code: "InvalidOptionParameter"}
	UnknownOption          = &Error{Code: "UnknownOption"}

	// usi
	InvalidPieceId = &Error{Code: "InvalidPieceId"}
)
