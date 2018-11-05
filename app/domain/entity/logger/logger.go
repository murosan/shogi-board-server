// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"os"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/config"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}

type logger struct {
	conf config.Config
}

func NewLogger(c config.Config) Logger {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	// TODO: confで
	//logrus.SetOutput(&lumberjack.Logger{
	//	Filename:  "log/application.log",
	//	MaxAge:    30,
	//	LocalTime: true,
	//	Compress:  true,
	//})
	return &logger{}
}

func (l *logger) Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func (l *logger) Info(args ...interface{}) {
	logrus.Info(args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func (l *logger) Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func (l *logger) Error(args ...interface{}) {
	logrus.Error(args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}
