// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exception

import (
	"github.com/pkg/errors"
)

var (
	// ErrTimeout means the execution exceeds them max time.
	ErrTimeout = errors.New("Execution timed out")
)
