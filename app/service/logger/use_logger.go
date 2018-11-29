// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/murosan/shogi-proxy-server/app/domain/logger"
	"github.com/murosan/shogi-proxy-server/app/service/config"
)

var l logger.Log = nil

func Use() logger.Log {
	if l == nil {
		l = logger.NewLogger(config.UseConfig())
	}
	return l
}
