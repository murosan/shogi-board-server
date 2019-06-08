// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package logger provides Logger interface and generate method.
package logger

import (
	"go.uber.org/zap"

	"github.com/murosan/shogi-board-server/app/config"
)

// Logger is a interface of logger.
type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}

// New generates zap.Config from given config,
// and returns it as a Logger instance.
func New(c *config.Config) Logger {
	l, err := c.Log.Build()

	if err != nil {
		panic(err)
	}

	return l
}
