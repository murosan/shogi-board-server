// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exception

import (
	"github.com/pkg/errors"
)

// FailedToClose wraps given error with message
// that means failed to close the shogi engine.
func FailedToClose(err error) error {
	return errors.Wrap(err, "Failed to close shogi engine")
}

var (
	// ErrTimeout means the execution exceeds them max time.
	ErrTimeout = errors.New("Execution timed out")
)

//var (
//	// engine
//
//	// ConnectionTimeout means got no response from shogi engine,
//	// or took too long time.
//	ConnectionTimeout = errors.New("ConnectionTimeout")
//
//	// FailedToConnect means failed to connect to shogi engine.
//	FailedToConnect = errors.New("FailedToConnect")
//
//	// FailedToClose means failed to disconnect to shogi engine.
//	FailedToClose = errors.New("FailedToClose")
//
//	// FailedToExecUSI means failed to execute USI command to shogi engine.
//	FailedToExecUSI = errors.New("FailedToExecUSI")
//
//	// usi
//
//	// UnknownPieceID means given piece ID is unknown.
//	UnknownPieceID = errors.New("UnknownPieceID")
//
//	// RowNumberIsOutOfRange means given row number is out of range.
//	RowNumberIsOutOfRange = errors.New("RowNumberIsOutOfRange")
//
//	// ColumnNumberIsOutOfRange given column number is out of range.
//	ColumnNumberIsOutOfRange = errors.New("ColumnNumberIsOutOfRange")
//
//	// UnknownCharacter means cannot parse the character.
//	UnknownCharacter = errors.New("UnknownCharacter")
//
//	// FailedToParseInfo means failed to parse usi info.
//	FailedToParseInfo = errors.New("FailedToParseInfo")
//)
