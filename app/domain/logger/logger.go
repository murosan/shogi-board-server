// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/murosan/shogi-board-server/app/domain/config"
	"go.uber.org/zap"
)

// Log ロギングライブラリ、zap のラッパー
type Log interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}

type log struct {
	z *zap.Logger
}

// NewLogger 新しい Log を返す
func NewLogger(c config.Config) Log {
	l, _ := c.GetLogConf().Build()
	return &log{l}
}

func (l *log) Debug(msg string, fields ...zap.Field) {
	l.z.Debug(msg, fields...)
}

func (l *log) Info(msg string, fields ...zap.Field) {
	l.z.Info(msg, fields...)
}

func (l *log) Warn(msg string, fields ...zap.Field) {
	l.z.Warn(msg, fields...)
}

func (l *log) Error(msg string, fields ...zap.Field) {
	l.z.Error(msg, fields...)
}

func (l *log) Fatal(msg string, fields ...zap.Field) {
	l.z.Fatal(msg, fields...)
}
