// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usi

import "sync"

// Result is a thought result of the shogi engine.
type Result struct {
	sync.RWMutex
	m map[int]*Info
}

// NewResult returns new Result
func NewResult() *Result {
	return &Result{
		m: make(map[int]*Info),
	}
}

// Set set the info value to map
func (r *Result) Set(key int, value *Info) {
	r.Lock()
	r.m[key] = value
	r.Unlock()
}

// GetAll returns all of the results.
func (r *Result) GetAll() map[int]*Info {
	r.RLock()
	defer r.RUnlock()
	// to avoid concurrent read and write
	m := make(map[int]*Info)
	for k, v := range r.m {
		m[k] = v
	}
	return m
}

// Flush cleans all result and set a new map
func (r *Result) Flush() {
	r.Lock()
	r.m = make(map[int]*Info)
	r.Unlock()
}
