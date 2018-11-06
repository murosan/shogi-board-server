// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"os"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/config"
	"github.com/sirupsen/logrus"
)

type Log interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

type log struct {
	conf config.Config
}

func NewLogger(c config.Config) Log {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	// TODO: confで
	//logrus.SetOutput(&lumberjack.Log{
	//	Filename:  "log/application.log",
	//	MaxAge:    30,
	//	LocalTime: true,
	//	Compress:  true,
	//})
	return &log{}
}

func (l *log) Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func (l *log) Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func (l *log) Info(args ...interface{}) {
	logrus.Info(args...)
}

func (l *log) Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func (l *log) Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func (l *log) Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func (l *log) Error(args ...interface{}) {
	logrus.Error(args...)
}

func (l *log) Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func (l *log) Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func (l *log) Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}
