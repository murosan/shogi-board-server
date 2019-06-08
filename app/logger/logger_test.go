// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"testing"

	"go.uber.org/zap"

	"github.com/murosan/goutils/testutils"
	"github.com/murosan/shogi-board-server/app/config"
)

func TestNew(t *testing.T) {
	c1 := &config.Config{}

	testutils.MustPanic(t, func() { New(c1) }, func(t *testing.T) {
		t.Helper()
		t.Error("[app > logger > New] Expected panic but there wasn't")
	})

	c2 := &config.Config{
		Log: zap.NewDevelopmentConfig(),
	}

	logger := New(c2)
	switch logger.(type) {
	case Logger:
		// success
	default:
		t.Error("[app > logger > New] Generated logger was not instanceof Logger")
	}
}
