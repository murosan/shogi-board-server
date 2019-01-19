// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package command

import "io"

// OsCmd exec.Command のラッパー
type OsCmd interface {
	Start() error
	Wait() error
	Write([]byte) error
	GetStdoutPipe() *io.ReadCloser
}
