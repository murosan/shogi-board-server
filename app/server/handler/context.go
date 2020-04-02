// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"net/url"

	"github.com/labstack/echo/v4"

	"github.com/murosan/shogi-board-server/app/logger"
)

// Context is a server contexts. Just wraps echo.Context.
type Context struct {
	ec     echo.Context
	logger logger.Logger
}

// NewContext returns new Context.
func NewContext(ec echo.Context, logger logger.Logger) *Context {
	return &Context{ec: ec, logger: logger}
}

func (ctx *Context) QueryParams() url.Values { return ctx.ec.QueryParams() }

func (ctx *Context) GetQuery(key string) string { return ctx.ec.QueryParam(key) }

func (ctx *Context) Bind(i interface{}) error { return ctx.ec.Bind(i) }

func (ctx *Context) NoContent(status int) error { return ctx.ec.NoContent(status) }

func (ctx *Context) JSON(status int, v interface{}) error { return ctx.ec.JSON(status, v) }
