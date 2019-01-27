// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

// NewResult returns new Result
func NewResult() *Result {
	return &Result{
		Result: make(map[int32]*Info),
	}
}

// NewInfo returns new Info
func NewInfo() *Info {
	return &Info{
		Values: make(map[string]int32),
	}
}

// NewMove returns new Move
func NewMove(source, dest *Point, pieceID int32, isPromoted bool) *Move {
	return &Move{
		Source:     source,
		Dest:       dest,
		PieceID:    pieceID,
		IsPromoted: isPromoted,
	}
}

// NewPoint returns new Point
func NewPoint(row, column int32) *Point {
	return &Point{
		Row:    row,
		Column: column,
	}
}

// NewButton returns new Button
func NewButton(name string) *Button {
	return &Button{Name: name}
}

// NewCheck returns new Check
func NewCheck(name string, val, init bool) *Check {
	return &Check{
		Name:    name,
		Val:     val,
		Default: init,
	}
}

// NewSpin returns new Spin
func NewSpin(name string, val, init, min, max int32) *Spin {
	return &Spin{
		Name:    name,
		Val:     val,
		Default: init,
		Min:     min,
		Max:     max,
	}
}

// NewSelect returns new Select
func NewSelect(name, val, init string, vars []string) *Select {
	return &Select{
		Name:    name,
		Val:     val,
		Default: init,
		Vars:    vars,
	}
}

// NewString returns new String
func NewString(name, val, init string) *String {
	return &String{
		Name:    name,
		Val:     val,
		Default: init,
	}
}

// NewFilename returns new Filename
func NewFilename(name, val, init string) *Filename {
	return &Filename{
		Name:    name,
		Val:     val,
		Default: init,
	}
}

// NewOptions returns new Options
func NewOptions() *Options {
	return &Options{
		Buttons:   make(map[string]*Button),
		Checks:    make(map[string]*Check),
		Spins:     make(map[string]*Spin),
		Selects:   make(map[string]*Select),
		Strings:   make(map[string]*String),
		Filenames: make(map[string]*Filename),
	}
}

// NewResponse returns new Response
func NewResponse() *Response {
	return &Response{}
}
