// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

type Connector interface {
	Connect() error
	Close() error
	Exec([]byte) error
	CatchEngineOutput()
}
